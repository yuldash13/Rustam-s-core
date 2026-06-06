package entity

import "errors"

var (
	EmptyID         = errors.New("id is empty")
	EmptyName       = errors.New("name is empty")
	NotFound        = errors.New("item not found")
	InvalidArgument = errors.New("invalid argument")
	Exist           = errors.New("item exist")
)
