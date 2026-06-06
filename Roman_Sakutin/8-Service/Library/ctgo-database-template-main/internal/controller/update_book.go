package controller

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	generated "github.com/project/library/generated/api/library"
	"github.com/project/library/internal/entity"
)

func (s *Service) UpdateBook(ctx context.Context, req *generated.UpdateBookRequest) (*generated.UpdateBookResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}

	bookToChange := entity.Book{
		ID:       req.GetId(),
		Name:     req.GetName(),
		AuthorID: append([]string(nil), req.GetAuthorIds()...),
	}
	if bookToChange.ID == "" {
		return nil, status.Error(codes.InvalidArgument, "book_id is empty")
	}
	if bookToChange.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "book_name is empty")
	}

	err := s.booksLibrary.UpdateBook(ctx, bookToChange)
	if err != nil {
		switch {
		case errors.Is(err, entity.InvalidArgument):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, entity.NotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, "failed to update book")
		}
	}

	return &generated.UpdateBookResponse{}, nil
}
