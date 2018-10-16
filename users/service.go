package users

import (
	"github.com/avasapollo/users-api/global"
	"github.com/sirupsen/logrus"
)

type service struct {
	logger   *logrus.Entry
	config   *global.AppConfig
	database Database
}

func NewService(le *logrus.Entry, conf *global.AppConfig, database Database) Service {
	return &service{
		logger:   le,
		config:   conf,
		database: database,
	}
}

func (svc service) CreateUser(user *User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	return svc.database.CreateUser(user)
}

func (svc service) GetUser(userID string) (*User, error) {
	return svc.database.GetUser(userID)
}

func (svc service) UpdateUser(user *User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	return svc.database.UpdateUser(user)
}

func (svc service) DeleteUser(userID string) error {
	return svc.database.DeleteUser(userID)
}
