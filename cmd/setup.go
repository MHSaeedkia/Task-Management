package cmd

import (
	"Task-Management/config"
	"Task-Management/internal/gateway"
	"log"

	"github.com/spf13/cobra"
)

var configPath string

func Run() error {
	rootCmd := &cobra.Command{
		Use: "Task-Manager",
		Run: func(cmd *cobra.Command, args []string) {
			// load config
			configs, err := config.LoadConfig(configPath)
			if err != nil {
				log.Fatal(err)
			}
			gateway.Run(configs)
		},
	}

	rootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "./config.yml", "config path ./config.yml")
	return rootCmd.Execute()
}
