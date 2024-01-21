package core

import "errors"

var (
	ErrUserNotFound      = errors.New("user doesn't exists")
	ErrSiteNotFound      = errors.New("site doesn't exists")
	ErrSiteAccessDenied  = errors.New("site access denied")
	ErrUserAlreadyExists = errors.New("user with such username already exists")
	ErrUnknown           = errors.New("unknown error")
)
