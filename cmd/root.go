package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/sultanowskii/godzilla/internal/cmd/server"
	"github.com/sultanowskii/godzilla/internal/logging"
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
		if err := logging.Setup(); err != nil {
			fmt.Println("error setting up the logger: ", err)
		}

		redisHost := config.Redis.Host
		redisPort := config.Redis.Port
		redisAddress := fmt.Sprintf("%s:%d", redisHost, redisPort)
		storage.InitRedisClient(redisAddress)

		redisConnErrorChan := make(chan error)
		defer close(redisConnErrorChan)
		go storage.MonitorRedisConnection(redisConnErrorChan)

		port, _ := cmd.Flags().GetInt16("port")
		address := fmt.Sprintf(":%d", port)

		e := server.Setup()

		serverErrChan := make(chan error)
		defer close(serverErrChan)
		go func(ch chan<- error) {
			err := e.Start(address)
			if err != nil {
				ch <- err
			}
		}(serverErrChan)

		for {
			select {
			case serverErr := <-serverErrChan:
				if serverErr != nil {
					logging.Fatal(fmt.Sprintf("server error: %s", serverErr.Error()))
					fmt.Printf("server error: %s", serverErr)
					return
				}
			case redisConnErr := <-redisConnErrorChan:
				if redisConnErr != nil {
					logging.Fatal(fmt.Sprintf("can't connect to redis: %s", redisConnErr))
					fmt.Printf("can't connect to redis: %s", redisConnErr)
					return
				}
			}
		}
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
