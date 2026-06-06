package controller

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	generated "github.com/project/library/generated/api/library"
	"github.com/project/library/internal/entity"
)

func (s *Service) GetAuthorBooks(req *generated.GetAuthorBooksRequest, stream generated.Library_GetAuthorBooksServer) error {
	if req == nil {
		return status.Error(codes.InvalidArgument, "request is nil")
	}

	authorID := req.GetAuthorId()
	if authorID == "" {
		return status.Error(codes.InvalidArgument, "author_id is empty")
	}

	books, err := s.booksLibrary.GetBookByAuthorID(stream.Context(), authorID)
	if err != nil {
		switch {
		case errors.Is(err, entity.InvalidArgument):
			return status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, entity.NotFound):
			return status.Error(codes.NotFound, err.Error())
		default:
			return status.Error(codes.Internal, "failed to get author books")
		}
	}

	for _, b := range books {
		if err := stream.Send(EntityBookToProto(b)); err != nil {
			if errors.Is(stream.Context().Err(), context.Canceled) {
				return status.Error(codes.Canceled, "request canceled by client")
			}
			if errors.Is(stream.Context().Err(), context.DeadlineExceeded) {
				return status.Error(codes.DeadlineExceeded, "request deadline exceeded")
			}
			return status.Error(codes.Internal, "failed to send book to stream")
		}
	}

	return nil
}
