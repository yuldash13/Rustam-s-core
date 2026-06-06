package logic

import (
	"context"

	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/domain"
)

func (s *service) GetUsers(ctx context.Context, filters *domain.UsersFilter) ([]domain.User, error) {
	user, err := s.repo.Users.GetUsers(ctx, filters)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return user, nil
}

func (s *service) GetUserByID(ctx context.Context, id int64) (*domain.User, []domain.Account, error) {
	user, _, err := s.repo.Users.GetUserByID(ctx, id)
	accounts, err := s.repo.Accounts.GetAccounts(ctx, id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, nil, err
	}
	return user, accounts, nil
}

func (s *service) CreateUser(ctx context.Context, u *domain.User) (int64, error) {
	id, err := s.repo.Users.CreateUser(ctx, u)
	if err != nil {
		s.logger.Error(err.Error())
		return id, err
	}
	return id, nil
}

func (s *service) UpdateUser(ctx context.Context, u *domain.User) error {
	err := s.repo.Users.UpdateUser(ctx, u)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}
