package cmd

import (
	"../conf"
	"database/sql"
	"fmt"
	"github.com/spf13/cobra"
	"log"

	_ "github.com/lib/pq"
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

	dbConnInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.DB.Host,
		config.DB.Port,
		config.DB.User,
		config.DB.Password,
		config.DB.Database,
		config.DB.SSLMode,
	)

	db, err := sql.Open("postgres", dbConnInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Connected to PSQL database")
}
