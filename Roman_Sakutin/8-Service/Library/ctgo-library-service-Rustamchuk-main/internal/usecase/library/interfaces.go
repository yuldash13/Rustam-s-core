package library

import "github.com/project/library/internal/entity"

type Library interface {
	Authors
	Books
}

type Authors interface {
	AddAuthor(a entity.Author) (int64, error)
	GetAuthor(id int64) (*entity.Author, error)
	UpdateAuthor(a entity.Author) error
}

type Books interface {
	AddBook(b entity.Book) (int64, error)
	GetBook(id int64) (*entity.Book, error)
	UpdateBook(b entity.Book) error
	GetBookByAuthorID(id int64) ([]Books, error)
}
