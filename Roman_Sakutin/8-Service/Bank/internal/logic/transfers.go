package logic

import (
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/db"
	"context"
)

func (s *service) GetTransferByID(ctx context.Context, id int, limit int) ([]db.Transfer, error) {
	item, err := s.repo.Transfers.GetTransferByID(ctx, id, limit)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return item, nil
}

func (s *service) MakeTransfer(ctx context.Context, t *db.Transfer) (int, error) {
	from, err1 := s.repo.Accounts.GetAccountByCurrency(ctx, t.IDFrom, t.Currency)
	if err1 != nil {
		s.logger.Error(err1.Error())
		return 0, err1
	}

	from.Balance -= t.Value

	to, err2 := s.repo.Accounts.GetAccountByCurrency(ctx, t.IDTo, t.Currency)
	if err2 != nil {
		s.logger.Error(err2.Error())
		return 0, err2
	}

	to.Balance += t.Value

	id, err := s.repo.Transfers.MakeTransfer(ctx, t)
	if err != nil {
		s.logger.Error(err.Error())
		return 0, err
	}
	return id, nil
}

func (s *service) CancelTransfer(ctx context.Context, id int) (*db.Transfer, error) {
	transfer, err1 := s.repo.Transfers.CancelTransfer(ctx, id)
	if err1 != nil {
		s.logger.Error(err1.Error())
		return nil, err1
	}

	var newTransfer = db.Transfer{
		IDFrom:    transfer.IDTo,
		IDTo:      transfer.IDFrom,
		TransDate: "",
		Currency:  transfer.Currency,
		Value:     transfer.Value,
	}

	from, err2 := s.repo.Accounts.GetAccountByCurrency(ctx, newTransfer.IDFrom, newTransfer.Currency)
	if err2 != nil {
		s.logger.Error(err2.Error())
		return nil, err2
	}

	from.Balance -= newTransfer.Value

	to, err3 := s.repo.Accounts.GetAccountByCurrency(ctx, newTransfer.IDTo, newTransfer.Currency)
	if err3 != nil {
		s.logger.Error(err3.Error())
		return nil, err3
	}

	to.Balance += newTransfer.Value

	_, err := s.repo.Transfers.MakeTransfer(ctx, &newTransfer)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}

	return &newTransfer, nil
}
