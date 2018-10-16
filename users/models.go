package users

import (
	"fmt"
	"time"
)

type User struct {
	ID        string
	Nickname  string
	Name      string
	Surname   string
	CreatedAt time.Time
}

func NewUser(id, nickname, name, surname string, createdAt time.Time) *User {
	return &User{
		ID:        id,
		Nickname:  nickname,
		Name:      name,
		Surname:   surname,
		CreatedAt: createdAt,
	}
}

func (u User) Validate() error {
	if u.ID == "" {
		return fmt.Errorf("id is empty")
	}
	if u.Name == "" {
		return fmt.Errorf("name is empty")
	}
	if u.Nickname == "" {
		return fmt.Errorf("nickname is empty")
	}
	if u.CreatedAt.IsZero() {
		return fmt.Errorf("created at is zero")
	}
	return nil
}
