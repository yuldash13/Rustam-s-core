package logic

import (
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/db"
	"context"
)

func (s *service) GetAccounts(ctx context.Context, id int) ([]db.Account, error) {
	account, err := s.repo.Accounts.GetAccounts(ctx, id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return account, nil
}

func (s *service) GetAccountByCurrency(ctx context.Context, id int, currency string) (*db.Account, error) {
	account, err := s.repo.Accounts.GetAccountByCurrency(ctx, id, currency)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return account, nil
}

func (s *service) GetAccountByID(ctx context.Context, id int) (*db.Account, error) {
	account, err := s.repo.Accounts.GetAccountByID(ctx, id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return account, nil
}

func (s *service) CreateAccount(ctx context.Context, a *db.Account) (int, error) {
	account, err := s.repo.Accounts.CreateAccount(ctx, a)
	if err != nil {
		s.logger.Error(err.Error())
		return 0, err
	}
	return account, nil
}

func (s *service) DepAccount(ctx context.Context, dep int, id int) error {
	err := s.repo.Accounts.DepAccount(ctx, dep, id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}

func (s *service) DeleteAccount(ctx context.Context, id int) error {
	err := s.repo.Accounts.DeleteAccount(ctx, id)
	if err != nil {
		s.logger.Error(err.Error())
		return err
	}
	return nil
}
