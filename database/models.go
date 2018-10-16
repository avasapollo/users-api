package database

import "time"

type UserDto struct {
	ID        string
	Nickname  string
	Name      string
	Surname   string
	CreatedAt time.Time
}

func NewUserDto(id, nickname, name, surname string, createdAt time.Time) *UserDto {
	return &UserDto{
		ID:        id,
		Nickname:  nickname,
		Name:      name,
		Surname:   surname,
		CreatedAt: createdAt,
	}
}
