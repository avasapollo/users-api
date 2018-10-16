package database

import (
	"fmt"

	"github.com/avasapollo/users-api/global"
	"github.com/avasapollo/users-api/users"
	"github.com/sirupsen/logrus"
)

type memoryDB struct {
	logger *logrus.Entry
	config *global.AppConfig
	users  map[string]*UserDto
}

func NewMemoryDB(le *logrus.Entry, conf *global.AppConfig) Database {
	return &memoryDB{
		logger: le,
		config: conf,
		users:  map[string]*UserDto{},
	}
}

func (m *memoryDB) CreateUser(user *users.User) error {
	_, err := m.GetUser(user.ID)
	if err != UserNotFound {
		return fmt.Errorf("couldn't possible process the operation")
	}
	m.users[user.ID] = NewUserDto(user.ID, user.Nickname, user.Name, user.Surname, user.CreatedAt)
	return nil
}
func (m memoryDB) GetUser(userID string) (*users.User, error) {
	user := m.users[userID]
	if user == nil {
		return nil, UserNotFound
	}
	return users.NewUser(user.ID, user.Nickname, user.Name, user.Surname, user.CreatedAt), nil
}
func (m *memoryDB) UpdateUser(user *users.User) error {
	_, err := m.GetUser(user.ID)
	if err != nil && err != UserNotFound {
		return fmt.Errorf("couldn't possible process the operation")
	}
	m.users[user.ID] = NewUserDto(user.ID, user.Nickname, user.Name, user.Surname, user.CreatedAt)
	return nil
}

func (m *memoryDB) DeleteUser(userID string) error {
	_, err := m.GetUser(userID)
	if err != nil && err != UserNotFound {
		return fmt.Errorf("couldn't possible process the operation")
	}
	delete(m.users, userID)
	return nil
}
