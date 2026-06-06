package library

import (
	"go.uber.org/zap"

	"github.com/project/library/internal/usecase/repository"
)

var _ AuthorsLibrary = (*LibraryImpl)(nil)
var _ BooksLibrary = (*LibraryImpl)(nil)

type LibraryImpl struct {
	logger        *zap.Logger
	authorsRepo   repository.AuthorsRepository
	booksRepo     repository.BooksRepository
	transactor    repository.Transactor
	outboxRepo    repository.OutboxRepository
	outboxEnabled bool
}

func New(logger *zap.Logger,
	authorsRepo repository.AuthorsRepository,
	booksRepo repository.BooksRepository,
	transactor repository.Transactor,
	outboxRepo repository.OutboxRepository,
	outboxEnabled bool,
) *LibraryImpl {
	return &LibraryImpl{
		logger:        logger,
		authorsRepo:   authorsRepo,
		booksRepo:     booksRepo,
		transactor:    transactor,
		outboxRepo:    outboxRepo,
		outboxEnabled: outboxEnabled,
	}
}
