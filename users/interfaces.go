package users

type Service interface {
	CreateUser(user *User) error
	GetUser(userID string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(userID string) error
}

type Database interface {
	CreateUser(user *User) error
	GetUser(userID string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(userID string) error
}
