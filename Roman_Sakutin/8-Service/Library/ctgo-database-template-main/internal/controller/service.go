package controller

import (
	"go.uber.org/zap"

	generated "github.com/project/library/generated/api/library"
	"github.com/project/library/internal/usecase/library"
)

// Service реализует интерфейс generated.LibraryServer
// Это контроллер, который обрабатывает gRPC запросы и делегирует работу use case
type Service struct {
	generated.UnimplementedLibraryServer // встроить для forward compatibility
	
	logger         *zap.Logger
	authorsLibrary library.AuthorsLibrary
	booksLibrary   library.BooksLibrary
}

// New создает новый экземпляр контроллера
// Принимает:
//   - logger: логгер для записи событий
//   - authorsLibrary: use case для работы с авторами
//   - booksLibrary: use case для работы с книгами
func New(
	logger *zap.Logger,
	authorsLibrary library.AuthorsLibrary,
	booksLibrary library.BooksLibrary,
) *Service {
	return &Service{
		logger:         logger,
		authorsLibrary: authorsLibrary,
		booksLibrary:   booksLibrary,
	}
}
