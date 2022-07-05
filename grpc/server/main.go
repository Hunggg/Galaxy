package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	metronion "github.com/metrogalaxy-io/metrogalaxy-api/grpc/proto/metronion/v1"
	wearables "github.com/metrogalaxy-io/metrogalaxy-api/grpc/proto/wearables/v1"
	"github.com/metrogalaxy-io/metrogalaxy-api/inject"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func RunGRPCServer(injector *inject.Injector, ctx context.Context) error {
	l := zap.S()
	s := grpc.NewServer()

	metronionController := NewMetronionController(injector)
	metronion.RegisterMetronionServiceServer(s, metronionController)
	wearablesController := NewWearablesController(injector)
	wearables.RegisterWearablesServiceServer(s, wearablesController)

	reflection.Register(s)

	port := viper.GetString("METROGALAXY_API_GRPC_PORT")
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		l.Errorw("listen error", "error", err)
		return err
	}
	l.Infof("gRPC Server is listening on port %v", port)

	go func() {
		<-ctx.Done()
		l.Infow("gracefully stop grpc server...")
		s.GracefulStop()
	}()

	return s.Serve(listener)
}

func RunHTTPServer(ctx context.Context) error {
	l := zap.S()
	engine := gin.Default()
	deployTime := time.Now()
	// gin middleware
	engine.Use(UseGinCORS())

	engine.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"timestamp":  time.Now().Unix(),
			"now":        time.Now(),
			"deployTime": deployTime,
			"version":    "0.0.1",
		})
	})

	gwmux := runtime.NewServeMux()
	port := viper.GetString("METROGALAXY_API_GRPC_PORT")
	proxyPort := viper.GetString("METROGALAXY_API_HTTP_PORT")
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("0.0.0.0:%v", port),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		l.Errorw("create grpc gateway proxies error")
	}
	if err := metronion.RegisterMetronionServiceHandler(ctx, gwmux, conn); err != nil {
		l.Errorw("register metronion http server handler error", "error", err)
		return err
	}
	if err := wearables.RegisterWearablesServiceHandler(ctx, gwmux, conn); err != nil {
		l.Errorw("register wearables http server handler error", "error", err)
		return err
	}

	engine.Any("/api/v1/*any", gin.WrapF(gwmux.ServeHTTP))

	httpServer := http.Server{
		Addr:    ":" + proxyPort,
		Handler: engine,
	}

	go func() {
		<-ctx.Done()
		l.Infow("gracefully stop http server...")
		_ = httpServer.Shutdown(ctx)
	}()

	l.Infof("gRPC Gateway is listening on port %v", proxyPort)
	return httpServer.ListenAndServe()
}
