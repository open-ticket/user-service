package main

import (
	"log"

	"github.com/open-ticket/user-service/cmd"
)

func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
