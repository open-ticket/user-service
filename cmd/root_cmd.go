package cmd

import (
	"fmt"
	"github.com/open-ticket/user-service/conf"
	"github.com/open-ticket/user-service/server"
	"github.com/spf13/cobra"
	"log"
)

func RootCommand() *cobra.Command {
	rootCmd := cobra.Command{
		Use: "start server",
		Run: run,
	}

	rootCmd.Flags().IntP("port", "p", 8000, "the port to run the server on")

	return &rootCmd
}

func run(cmd *cobra.Command, args []string) {
	config, err := conf.LoadConfig(cmd)
	if err != nil {
		log.Fatal("Error loading config: " + err.Error())
	}

	fmt.Println("Starting user-service...")

	server.StartServer(config)
}
