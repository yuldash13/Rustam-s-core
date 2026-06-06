package logic

import (
	"context"

	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/domain"
)

func (s *service) GetAccounts(ctx context.Context, IDUser int64) ([]domain.Account, error) {
	account, err := s.repo.Accounts.GetAccounts(ctx, IDUser)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return account, nil
}

func (s *service) GetAccountByCurrency(ctx context.Context, id int64, currency string) (*domain.Account, error) {
	account, err := s.repo.Accounts.GetAccountByCurrency(ctx, id, currency)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return account, nil
}

func (s *service) GetAccountByID(ctx context.Context, id int64) (*domain.Account, error) {
	account, err := s.repo.Accounts.GetAccountByID(ctx, id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return account, nil
}

func (s *service) CreateAccount(ctx context.Context, a *domain.Account) (int64, error) {
	account, err := s.repo.Accounts.CreateAccount(ctx, a)
	if err != nil {
		s.logger.Error(err.Error())
		return 0, err
	}
	return account, nil
}

func (s *service) UpdateAccount(ctx context.Context, id int64, balance int64) error {
	err := s.repo.Accounts.UpdateAccount(ctx, id, balance)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *service) DepAccount(ctx context.Context, dep domain.DepAccount, id int64) error {
	err := s.repo.Accounts.DepAccount(ctx, dep, id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *service) DeleteAccount(ctx context.Context, id int64) error {
	err := s.repo.Accounts.DeleteAccount(ctx, id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}
