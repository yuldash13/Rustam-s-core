package controller

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	generated "github.com/project/library/generated/api/library"
	"github.com/project/library/internal/entity"
)

var tracer = otel.Tracer("library-service")

func (s *Service) AddBook(ctx context.Context, req *generated.AddBookRequest) (*generated.AddBookResponse, error) {
	ctx, span := tracer.Start(ctx, "HandleAddBook")
	defer span.End()

	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}
	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "book_name is empty")
	}

	bookToCreate := entity.Book{
		Name:     req.GetName(),
		AuthorID: append([]string(nil), req.GetAuthorIds()...),
	}

	createdBook, err := s.booksLibrary.CreateBook(ctx, bookToCreate)
	if err != nil {
		switch {
		case errors.Is(err, entity.InvalidArgument), errors.Is(err, entity.EmptyName):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, entity.NotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.Is(err, entity.Exist):
			return nil, status.Error(codes.AlreadyExists, err.Error())
		default:
			return nil, status.Error(codes.Internal, "failed to create book")
		}
	}

	return &generated.AddBookResponse{
		Book: EntityBookToProto(createdBook),
	}, nil
}
