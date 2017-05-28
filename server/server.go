package server

import (
	"fmt"
	"github.com/open-ticket/user-service/conf"
	database "github.com/open-ticket/user-service/db"
	"log"
	"net/http"
)

func StartServer(config *conf.Config) {
	portString := fmt.Sprintf(":%d", config.Port)

	db := database.Connect(config)
	defer db.Close()

	log.Fatal(http.ListenAndServe(portString, nil))
}
