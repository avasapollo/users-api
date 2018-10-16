package database

import (
	"github.com/avasapollo/users-api/users"
)

type Database interface {
	CreateUser(user *users.User) error
	GetUser(userID string) (*users.User, error)
	UpdateUser(user *users.User) error
	DeleteUser(userID string) error
}
