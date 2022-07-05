package cmd

import (
	"context"
	"os"
	"os/signal"

	grpcServer "github.com/metrogalaxy-io/metrogalaxy-api/grpc/server"
	"github.com/metrogalaxy-io/metrogalaxy-api/inject"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start http server with configured api",
	Long:  `Starts a http server and serves the configured api`,
	RunE: func(cmd *cobra.Command, args []string) error {
		l := zap.S()
		// initialize injectors
		injector := inject.NewInjector()
		if err := injector.Run(context.Background()); err != nil {
			return err
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		go func() {
			l.Fatal(grpcServer.RunGRPCServer(injector, ctx))
		}()

		go func() {
			l.Fatal(grpcServer.RunHTTPServer(ctx))
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		sig := <-quit
		cancel()
		l.Infof("Shutting down server... Reason: %v", sig)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
	viper.SetDefault("METROGALAXY_API_GRPC_PORT", 9000)
	viper.SetDefault("METROGALAXY_API_HTTP_PORT", 8080)
}
