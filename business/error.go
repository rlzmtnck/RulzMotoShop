package business

import "errors"

var (
	ErrInternalServer = errors.New("500 Internal Server Error")
	ErrEmailorPass    = errors.New("Email or Password is incorrect")
	ErrDuplicateData  = errors.New("Account already exist")
	ErrNotFound       = errors.New("404 Not Found")
	ErrUnathorized    = errors.New("Unauthorized - Go Away")
)
