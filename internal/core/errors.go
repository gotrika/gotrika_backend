package core

import "errors"

var (
	ErrUserNotFound      = errors.New("user doesn't exists")
	ErrUserAlreadyExists = errors.New("user with such username already exists")
	ErrUnknown           = errors.New("unknown error")
)
