package controller

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	generated "github.com/project/library/generated/api/library"
	"github.com/project/library/internal/entity"
)

func (s *Service) GetBookInfo(ctx context.Context, req *generated.GetBookInfoRequest) (*generated.GetBookInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}

	bookID := req.GetId()
	if bookID == "" {
		return nil, status.Error(codes.InvalidArgument, "book_id is empty")
	}

	book, err := s.booksLibrary.GetBook(ctx, bookID)
	if err != nil {
		switch {
		case errors.Is(err, entity.InvalidArgument):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, entity.NotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, "failed to get book")
		}
	}

	return &generated.GetBookInfoResponse{
		Book: EntityBookToProto(book),
	}, nil
}
