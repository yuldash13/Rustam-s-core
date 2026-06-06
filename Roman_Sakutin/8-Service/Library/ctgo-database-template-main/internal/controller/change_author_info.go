package controller

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	generated "github.com/project/library/generated/api/library"
	"github.com/project/library/internal/entity"
)

func (s *Service) ChangeAuthorInfo(ctx context.Context, req *generated.ChangeAuthorInfoRequest) (*generated.ChangeAuthorInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}

	authorToChange := entity.Author{
		ID:   req.GetId(),
		Name: req.GetName(),
	}
	if authorToChange.ID == "" || authorToChange.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "author_id is empty")
	}
	if authorToChange.Name == "" {
		return nil, status.Error(codes.InvalidArgument, "author_name is empty")
	}

	err := s.authorsLibrary.UpdateAuthor(ctx, authorToChange)
	if err != nil {
		switch {
		case errors.Is(err, entity.InvalidArgument):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, entity.NotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, "failed to change author")
		}
	}

	return &generated.ChangeAuthorInfoResponse{}, nil
}
