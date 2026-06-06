package logic

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/model/domain"
	"Roman_Sakutin/Roman_Sakutin/8-Service/Bank/internal/repo"
)

// setupTestService создает тестовый сервис с моками
func setupTestService(t *testing.T) (*service, *repo.MockAccounts, *gomock.Controller) {
	ctrl := gomock.NewController(t)
	mockAccounts := repo.NewMockAccounts(ctrl)

	// Создаем мок репозитория
	mockRepo := &repo.Repo{
		Accounts: mockAccounts,
	}

	// Создаем logger (можно использовать noop logger для тестов)
	logger := zap.NewNop()

	serv := &service{
		logger: logger,
		repo:   mockRepo,
	}

	return serv, mockAccounts, ctrl
}

// TestGetAccounts - табличные тесты для GetAccounts с использованием require
func TestGetAccounts(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		name           string
		userID         int64
		mockAccounts   []domain.Account
		mockError      error
		expectedCount  int
		expectError    bool
		expectedError  error
		validateResult func(t *testing.T, accounts []domain.Account)
	}{
		{
			name:   "Успешное получение нескольких счетов",
			userID: 1,
			mockAccounts: []domain.Account{
				{ID: 1, IDUser: 1, Balance: 1000, Currency: "USD"},
				{ID: 2, IDUser: 1, Balance: 500, Currency: "EUR"},
			},
			mockError:     nil,
			expectedCount: 2,
			expectError:   false,
			validateResult: func(t *testing.T, accounts []domain.Account) {
				require.Len(t, accounts, 2)
				require.Equal(t, int64(1), accounts[0].ID)
				require.Equal(t, "USD", accounts[0].Currency)
				require.Equal(t, int64(1000), accounts[0].Balance)
				require.Equal(t, int64(2), accounts[1].ID)
				require.Equal(t, "EUR", accounts[1].Currency)
			},
		},
		{
			name:          "Успешное получение одного счета",
			userID:        2,
			mockAccounts:  []domain.Account{{ID: 10, IDUser: 2, Balance: 2500, Currency: "RUB"}},
			mockError:     nil,
			expectedCount: 1,
			expectError:   false,
			validateResult: func(t *testing.T, accounts []domain.Account) {
				require.Len(t, accounts, 1)
				require.Equal(t, int64(10), accounts[0].ID)
				require.Equal(t, int64(2500), accounts[0].Balance)
				require.Equal(t, "RUB", accounts[0].Currency)
			},
		},
		{
			name:          "Пустой список счетов",
			userID:        999,
			mockAccounts:  []domain.Account{},
			mockError:     nil,
			expectedCount: 0,
			expectError:   false,
			validateResult: func(t *testing.T, accounts []domain.Account) {
				require.NotNil(t, accounts)
				require.Len(t, accounts, 0)
			},
		},
		{
			name:          "Ошибка базы данных",
			userID:        3,
			mockAccounts:  nil,
			mockError:     errors.New("database connection error"),
			expectedCount: 0,
			expectError:   true,
			expectedError: errors.New("database connection error"),
			validateResult: func(t *testing.T, accounts []domain.Account) {
				require.Nil(t, accounts)
			},
		},
		{
			name:          "ID пользователя равен 0",
			userID:        0,
			mockAccounts:  []domain.Account{},
			mockError:     nil,
			expectedCount: 0,
			expectError:   false,
			validateResult: func(t *testing.T, accounts []domain.Account) {
				require.NotNil(t, accounts)
				require.Len(t, accounts, 0)
			},
		},
		{
			name:          "Отрицательный ID пользователя",
			userID:        -1,
			mockAccounts:  []domain.Account{},
			mockError:     nil,
			expectedCount: 0,
			expectError:   false,
			validateResult: func(t *testing.T, accounts []domain.Account) {
				require.NotNil(t, accounts)
				require.Len(t, accounts, 0)
			},
		},
		{
			name:          "Ошибка контекста",
			userID:        1,
			mockAccounts:  nil,
			mockError:     context.Canceled,
			expectedCount: 0,
			expectError:   true,
			expectedError: context.Canceled,
			validateResult: func(t *testing.T, accounts []domain.Account) {
				require.Nil(t, accounts)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			serv, mockAccounts, ctrl := setupTestService(t)
			defer ctrl.Finish()

			// Настраиваем мок
			mockAccounts.EXPECT().
				GetAccounts(ctx, tc.userID).
				Return(tc.mockAccounts, tc.mockError).
				Times(1)

			// Выполняем тест
			accounts, err := serv.GetAccounts(ctx, tc.userID)

			// Проверяем ошибку с помощью require
			if tc.expectError {
				require.Error(t, err)
				if tc.expectedError != nil {
					require.Equal(t, tc.expectedError, err)
				}
			} else {
				require.NoError(t, err)
			}

			// Проверяем количество счетов
			if !tc.expectError {
				require.NotNil(t, accounts)
				require.Len(t, accounts, tc.expectedCount)
			} else {
				require.Nil(t, accounts)
			}

			// Дополнительная валидация результата, если указана
			if tc.validateResult != nil {
				tc.validateResult(t, accounts)
			}
		})
	}
}
