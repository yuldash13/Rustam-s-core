package logic

import (
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/db"
	"context"
)

func (s *service) GetTransfers(ctx context.Context, id int, limit int) ([]db.Transfer, error) {
	item, err := s.repo.Transfers.GetTransfers(ctx, id, limit)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return item, nil
}

func (s *service) GetTransferByID(ctx context.Context, id int) (*db.Transfer, error) {
	item, err := s.repo.Transfers.GetTransferByID(ctx, id)
	if err != nil {
		s.logger.Error(err.Error())
		return nil, err
	}
	return item, nil
}

func (s *service) MakeTransfer(ctx context.Context, t *db.Transfer) (int, error) {
	var id int
	var err error

	err = s.repo.Transaction.PerformTransaction(ctx, func(ctx context.Context) error {
		from, err1 := s.repo.Accounts.GetAccountByCurrency(ctx, t.IDFrom, t.Currency)
		if err1 != nil {
			s.logger.Error(err1.Error())
			return err1
		}

		to, err2 := s.repo.Accounts.GetAccountByCurrency(ctx, t.IDTo, t.Currency)
		if err2 != nil {
			s.logger.Error(err2.Error())
			return err2
		}

		if from.Currency != to.Currency {
			return db.WrongCurrency
		}

		if from.Balance < t.Value {
			return db.NotEnough
		}

		from.Balance -= t.Value
		to.Balance += t.Value

		err3 := s.repo.Accounts.UpdateAccount(ctx, from.ID, from.Balance)
		if err3 != nil {
			s.logger.Error(err3.Error())
			return err3
		}
		err4 := s.repo.Accounts.UpdateAccount(ctx, to.ID, to.Balance)
		if err4 != nil {
			s.logger.Error(err4.Error())
			return err4
		}

		id, err = s.repo.Transfers.MakeTransfer(ctx, t)
		if err != nil {
			s.logger.Error(err.Error())
			return err
		}
		return nil
	})

	return id, err
}

func (s *service) CancelTransfer(ctx context.Context, id int) error {
	var err error

	err = s.repo.Transaction.PerformTransaction(ctx, func(ctx context.Context) error {
		transfer, err1 := s.repo.Transfers.GetTransferByID(ctx, id)
		if err1 != nil {
			s.logger.Error(err1.Error())
			return err1
		}
		err0 := s.repo.Transfers.CancelTransfer(ctx, id)
		if err0 != nil {
			s.logger.Error(err0.Error())
			return err0
		}

		var newTransfer = db.Transfer{
			IDFrom:   transfer.IDTo,
			IDTo:     transfer.IDFrom,
			Currency: transfer.Currency,
			Value:    transfer.Value,
		}

		_, err2 := s.MakeTransfer(ctx, &newTransfer)
		if err1 != nil {
			s.logger.Error(err2.Error())
			return err2
		}
		return nil
	})

	return err
}
