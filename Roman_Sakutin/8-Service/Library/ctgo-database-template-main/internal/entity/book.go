package entity

import "time"

type Book struct {
	ID        string
	Name      string
	AuthorID  []string
	CreatedAt time.Time
	UpdatedAt time.Time
}
