package server

import (
	"database/sql"
	"log"

	"github.com/open-ticket/user-service/models/user"
)

type UserService interface {
	getAllUsers() (string, error)
	getUser(userId int) (string, error)
	updateUser(userId int) (string, error)
}

type userService struct {
	db *sql.DB
}

func (svc *userService) getAllUsers() ([]user.User, error) {
	log.Println("hello")
	users, err := user.Find(svc.db)
	return users, err
}

func (svc *userService) getUser(userId int) (string, error) {
	return "test", nil
}

func createUserService(db *sql.DB) userService {
	svc := userService{db}
	return svc
}
