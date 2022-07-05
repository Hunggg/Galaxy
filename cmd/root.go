package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	// "github.com/joho/godotenv"

	"github.com/metrogalaxy-io/metrogalaxy-api/env"
	"github.com/metrogalaxy-io/metrogalaxy-api/libs/logging"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "metrogalaxy-api",
	Short: "MetroGalaxy API",
	Long:  `MetroGalaxy API`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// RunE: func(cmd *cobra.Command, args []string) error {

	// },
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config.yaml)")
	// RootCmd.PersistentFlags().Bool("db_debug", false, "log sql to console")
	// viper.BindPFlag("db_debug", RootCmd.PersistentFlags().Lookup("db_debug"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}

		_, b, _, _ := runtime.Caller(0)
		basepath := filepath.Dir(b)

		// Search config in home directory with name .config (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath("./env/")
		viper.AddConfigPath("../env/")
		viper.AddConfigPath("../../env/")
		viper.AddConfigPath(basepath)
		viper.SetConfigType("json")

		network := os.Getenv("NETWORK")
		log.Printf("On network: %v", network)
		switch network {
		case "avax":
			viper.SetConfigName("avax")
		case "fuji":
			viper.SetConfigName("fuji")
		default:
			viper.SetConfigName("hardhat")
		}
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Printf("Using config file: %v\n", viper.ConfigFileUsed())
	} else {
		log.Fatalf("error read config: %v\n", err.Error())
	}

	cfg := env.Config{}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("error parse config file: %v\n", err.Error())
	}

	env.InitConfig(&cfg)

	l, flush, err := logging.NewSugaredLogger()
	if err != nil {
		panic(err)
	}

	defer func() {
		flush()
	}()
	zap.ReplaceGlobals(l.Desugar())
}
