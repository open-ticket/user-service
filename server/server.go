package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/open-ticket/user-service/conf"
	database "github.com/open-ticket/user-service/db"
)

func StartServer(config *conf.Config) {
	portString := fmt.Sprintf(":%d", config.Port)

	db := database.Connect(config)
	defer db.Close()

	svc := createUserService(db)
	handlers := setupEndpoints(svc)
	for _, handler := range handlers {
		http.Handle(handler.path, handler.handler)
	}

	log.Println("Starting server")

	log.Fatal(http.ListenAndServe(portString, nil))
}
