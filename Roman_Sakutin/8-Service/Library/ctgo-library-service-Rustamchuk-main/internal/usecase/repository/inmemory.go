package repository

import (
	"sync"

	"github.com/project/library/internal/entity"
)

type AuthorsRepo struct {
	DB   map[int64]entity.Author
	RWMu sync.RWMutex
}

func NewAuthorsRepo() *AuthorsRepo {
	return &AuthorsRepo{
		DB:   make(map[int64]entity.Author, 1000),
		RWMu: sync.RWMutex{},
	}
}

func (AR *AuthorsRepo) AddAuthor(a entity.Author) (int64, error) {
	_, exist := AR.DB[a.ID]
	if exist {
		return a.ID, entity.Exist
	}

	AR.RWMu.RLock()
	AR.DB[a.ID] = a
	AR.RWMu.RUnlock()
	return a.ID, nil
}

func (AR *AuthorsRepo) GetAuthor(id int64) (*entity.Author, error) {
	value, exist := AR.DB[id]
	if !exist {
		return nil, entity.NotFound
	}
	return &value, nil
}

func (AR *AuthorsRepo) UpdateAuthor(a entity.Author) error {
	_, exist := AR.DB[a.ID]
	if !exist {
		return entity.NotFound
	}

	AR.RWMu.RLock()
	AR.DB[a.ID] = a
	AR.RWMu.RUnlock()
	return nil
}

type BooksRepo struct {
	DB   map[int64]entity.Book
	RWMu sync.RWMutex
}

func NewBooksRepo() *BooksRepo {
	return &BooksRepo{
		DB:   make(map[int64]entity.Book, 1000),
		RWMu: sync.RWMutex{},
	}
}

func (BR *BooksRepo) AddBook(b entity.Book) (int64, error) {
	_, exist := BR.DB[b.ID]
	if exist {
		return b.ID, entity.Exist
	}

	BR.RWMu.RLock()
	BR.DB[b.ID] = b
	BR.RWMu.RUnlock()
	return b.ID, nil
}

func (BR *BooksRepo) GetBook(id int64) (*entity.Book, error) {
	value, exist := BR.DB[id]
	if !exist {
		return nil, entity.NotFound
	}
	return &value, nil
}

func (BR *BooksRepo) UpdateBook(b entity.Book) error {
	_, exist := BR.DB[b.ID]
	if !exist {
		return entity.NotFound
	}

	BR.RWMu.RLock()
	BR.DB[b.ID] = b
	BR.RWMu.RUnlock()
	return nil
}

func (BR *BooksRepo) GetBookByAuthorID(id int64) ([]Books, error) {
	_, exist := BR.DB[id]
	if !exist {
		return nil, entity.NotFound
	}
	return nil, nil
}
