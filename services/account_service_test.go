package services

import (
	"context"
	"errors"
	"testing"

	"pulseledger/dto"
	"pulseledger/entities"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockAccountRepository struct {
	mock.Mock
}

func (m *mockAccountRepository) Create(ctx context.Context, account *entities.Account) error {
	args := m.Called(ctx, account)
	return args.Error(0)
}

func (m *mockAccountRepository) GetByID(ctx context.Context, id uint64) (*entities.Account, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entities.Account), args.Error(1)
}

func TestAccountService_CreateAccount_Success(t *testing.T) {
	mockRepo := new(mockAccountRepository)
	service := NewAccountService(mockRepo)
	req := dto.CreateAccountRequest{DocumentNumber: 12345678900}

	mockRepo.On("Create", mock.Anything, mock.AnythingOfType("*entities.Account")).Return(nil)

	err := service.CreateAccount(context.Background(), req)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestAccountService_CreateAccount_ValidationFail(t *testing.T) {
	service := NewAccountService(nil)
	req := dto.CreateAccountRequest{}

	err := service.CreateAccount(context.Background(), req)
	assert.Error(t, err)
}

func TestAccountService_GetAccountByID_Success(t *testing.T) {
	mockRepo := new(mockAccountRepository)
	service := NewAccountService(mockRepo)
	expected := &entities.Account{
		BaseEntity:     entities.BaseEntity{ID: 1},
		DocumentNumber: 12345678900,
	}
	req := dto.GetAccountByIDRequest{AccountID: 1}

	mockRepo.On("GetByID", mock.Anything, uint64(1)).Return(expected, nil)

	resp, err := service.GetAccountByID(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, expected.ID, resp.AccountId)
	assert.Equal(t, expected.DocumentNumber, resp.DocumentNumber)
	mockRepo.AssertExpectations(t)
}

func TestAccountService_GetAccountByID_ValidationFail(t *testing.T) {
	service := NewAccountService(nil)
	req := dto.GetAccountByIDRequest{}

	resp, err := service.GetAccountByID(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestAccountService_GetAccountByID_RepoError(t *testing.T) {
	mockRepo := new(mockAccountRepository)
	service := NewAccountService(mockRepo)
	req := dto.GetAccountByIDRequest{AccountID: 99}

	mockRepo.On("GetByID", mock.Anything, uint64(99)).Return(&entities.Account{}, errors.New("not found"))

	resp, err := service.GetAccountByID(context.Background(), req)
	assert.Error(t, err)
	assert.Nil(t, resp)
}
