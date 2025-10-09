package db

import "errors"

var (
	NotFound = errors.New("item not found")
	Exist    = errors.New("item exist")
)
