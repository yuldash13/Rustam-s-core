package db

import "errors"

var (
	NotFound      = errors.New("item not found")
	Exist         = errors.New("item exist")
	NotEnough     = errors.New("not enough")
	WrongCurrency = errors.New("wrong currency")
)
