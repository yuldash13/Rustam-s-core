package entity

import "errors"

var (
	NotFound        = errors.New("item not found")
	InvalidArgument = errors.New("invalid argument")
	Exist           = errors.New("item exist")
)
