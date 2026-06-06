package controller

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	generated "github.com/project/library/generated/api/library"
	"github.com/project/library/internal/entity"
)

func (s *Service) GetAuthorInfo(ctx context.Context, req *generated.GetAuthorInfoRequest) (*generated.GetAuthorInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request is nil")
	}

	authorID := req.GetId()
	if authorID == "" {
		return nil, status.Error(codes.InvalidArgument, "author_id is empty")
	}

	item, err := s.authorsLibrary.GetAuthor(ctx, authorID)
	if err != nil {
		switch {
		case errors.Is(err, entity.InvalidArgument):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		case errors.Is(err, entity.NotFound):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, "failed to get author")
		}
	}

	return &generated.GetAuthorInfoResponse{
		Id:   item.ID,
		Name: item.Name,
	}, nil
}
