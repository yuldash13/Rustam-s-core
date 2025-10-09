package logic

import (
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/db"
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/repo"
	"context"
	"go.uber.org/zap"
)

type Service interface {
	Users
	Accounts
	Transfers
}

type Users interface {
	GetUsers(ctx context.Context, filters *db.UsersFilter) ([]db.User, error)
	GetUserByID(ctx context.Context, id int) (*db.User, error)
	CreateUser(ctx context.Context, u *db.User) (int, error)
	UpdateUser(ctx context.Context, id int, u *db.User) error
}

type Accounts interface {
	GetAccounts(ctx context.Context, id int) ([]db.Account, error)
	GetAccountByCurrency(ctx context.Context, id int, currency string) (*db.Account, error)
	GetAccountByID(ctx context.Context, id int) (*db.Account, error)
	CreateAccount(ctx context.Context, a *db.Account) (int, error)
	DepAccount(ctx context.Context, dep int, id int) error
	DeleteAccount(ctx context.Context, id int) error
}

type Transfers interface {
	GetTransferByID(ctx context.Context, id int, limit int) ([]db.Transfer, error)
	MakeTransfer(ctx context.Context, t *db.Transfer) (int, error)
	CancelTransfer(ctx context.Context, id int) (*db.Transfer, error)
}

type service struct {
	logger *zap.Logger
	repo   *repo.Repo
}

func NewService(logger *zap.Logger, repo *repo.Repo) Service {
	return &service{logger: logger, repo: repo}
}
