package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // postgres driver
	"github.com/open-ticket/user-service/conf"
	"log"
)

func Connect(config *conf.Config) *sql.DB {
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

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Connected to PSQL database")
	return db
}
