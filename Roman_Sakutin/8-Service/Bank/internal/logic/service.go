package logic

import (
	"context"

	"go.uber.org/zap"

	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/domain"
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/repo"
)

type Service interface {
	Users
	Accounts
	Transfers
}

type Users interface {
	GetUsers(ctx context.Context, filters *domain.UsersFilter) ([]domain.User, error)
	GetUserByID(ctx context.Context, id int64) (*domain.User, []domain.Account, error)
	CreateUser(ctx context.Context, u *domain.User) (int64, error)
	UpdateUser(ctx context.Context, u *domain.User) error
}

type Accounts interface {
	GetAccounts(ctx context.Context, IDUser int64) ([]domain.Account, error)
	GetAccountByCurrency(ctx context.Context, id int64, currency string) (*domain.Account, error)
	GetAccountByID(ctx context.Context, id int64) (*domain.Account, error)
	CreateAccount(ctx context.Context, a *domain.Account) (int64, error)
	UpdateAccount(ctx context.Context, id int64, balance int64) error
	DepAccount(ctx context.Context, dep domain.DepAccount, id int64) error
	DeleteAccount(ctx context.Context, id int64) error
}

type Transfers interface {
	GetTransfers(ctx context.Context, id int64, limit int64) ([]domain.Transfer, error)
	GetTransferByID(ctx context.Context, id int64) (*domain.Transfer, error)
	MakeTransfer(ctx context.Context, t *domain.Transfer) (int64, error)
	CancelTransfer(ctx context.Context, id int64) error
}

type service struct {
	logger *zap.Logger
	repo   *repo.Repo
}

func NewService(logger *zap.Logger, repo *repo.Repo) Service {
	return &service{logger: logger, repo: repo}
}
