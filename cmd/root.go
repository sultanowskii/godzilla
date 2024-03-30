package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/sultanowskii/godzilla/internal/cmd/server"
	"github.com/sultanowskii/godzilla/pkg/storage"
)

var (
	configFile string
	config     Config
)

var rootCmd = &cobra.Command{
	Use:   "godzilla",
	Short: "URL shortener & custom URL creator",
	Long:  "Go, Dzilla",
	Run: func(cmd *cobra.Command, args []string) {
		redisHost := config.Redis.Host
		redisPort := config.Redis.Port
		redisAddress := fmt.Sprintf("%s:%d", redisHost, redisPort)
		storage.InitRedisClient(redisAddress)

		port, _ := cmd.Flags().GetInt16("port")
		address := fmt.Sprintf(":%d", port)

		e := server.SetupEcho()
		e.Logger.Info(e.Start(address))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&configFile, "config", "godzilla.yaml", "Config file (default is godzilla.yaml)")

	rootCmd.Flags().Int16P("port", "p", 8431, "Server port")
}

func initConfig() {
	if err := InitConfig(); err != nil {
		fmt.Printf("Error reading config: %s\n", err)
		return
	}

	if err := ParseConfig(&config); err != nil {
		fmt.Printf("Error reading config: %s\n", err)
		return
	}
}
