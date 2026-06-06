package controller

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	generated "github.com/project/library/generated/api/library"
	"github.com/project/library/internal/entity"
)

func (s *Service) RegisterAuthor(ctx context.Context, req *generated.RegisterAuthorRequest) (*generated.RegisterAuthorResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}

	authorToCreate := entity.Author{
		Name: req.GetName(),
	}

	createdAuthor, err := s.authorsLibrary.CreateAuthor(ctx, authorToCreate)
	if err != nil {
		switch {
		case errors.Is(err, entity.InvalidArgument):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, entity.Exist):
			return nil, status.Error(codes.AlreadyExists, err.Error())
		default:
			return nil, status.Error(codes.Internal, "failed to create author")
		}
	}

	return &generated.RegisterAuthorResponse{
		Id: createdAuthor.ID,
	}, nil
}
