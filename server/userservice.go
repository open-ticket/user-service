package server

type UserService interface {
	AllUsers() (string, error)
	User(userId int) (string, error)
}
