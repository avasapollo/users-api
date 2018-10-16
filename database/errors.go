package database

import "errors"

var UserNotFound = errors.New("user not found")
var UserAlreadyCreated = errors.New("user is already created")
