package services

import (
	"context"
	"errors"
	"testing"

	"pulseledger/dto"
	"pulseledger/entities"
	"pulseledger/enums"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockTransactionRepository struct {
	mock.Mock
}

func (m *mockTransactionRepository) Create(ctx context.Context, tx *entities.Transaction) error {
	args := m.Called(ctx, tx)
	return args.Error(0)
}

func TestTransactionService_CreateTransaction_Success(t *testing.T) {
	mockRepo := new(mockTransactionRepository)
	service := NewTransactionService(mockRepo)

	req := dto.CreateTransactionRequest{
		AccountID:       1,
		OperationTypeID: enums.NormalPurchase,
		Amount:          decimal.NewFromFloat(100.00),
	}

	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*entities.Transaction")).Return(nil)

	err := service.CreateTransaction(context.Background(), req)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestTransactionService_CreateTransaction_ValidationFail(t *testing.T) {
	service := NewTransactionService(nil)

	req := dto.CreateTransactionRequest{}

	err := service.CreateTransaction(context.Background(), req)

	assert.Error(t, err)
}

func TestTransactionService_CreateTransaction_RepoError(t *testing.T) {
	mockRepo := new(mockTransactionRepository)
	service := NewTransactionService(mockRepo)

	req := dto.CreateTransactionRequest{
		AccountID:       1,
		OperationTypeID: enums.NormalPurchase,
		Amount:          decimal.NewFromFloat(100.00),
	}

	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*entities.Transaction")).Return(errors.New("db error"))

	err := service.CreateTransaction(context.Background(), req)

	assert.Error(t, err)
	assert.Equal(t, "db error", err.Error())
	mockRepo.AssertExpectations(t)
}
