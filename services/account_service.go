package services

import (
	"context"
	"fmt"
	"pulseledger/dto"
	"pulseledger/entities"
	"pulseledger/repositories"
	"pulseledger/validator"

	log "github.com/sirupsen/logrus"
)

type AccountService interface {
	CreateAccount(ctx context.Context, req dto.CreateAccountRequest) error
	GetAccountByID(ctx context.Context, req dto.GetAccountByIDRequest) (*dto.GetAccountByIDResponse, error)
}

type accountService struct {
	repo repositories.AccountRepository
}

func NewAccountService(repo repositories.AccountRepository) AccountService {
	return &accountService{repo}
}

func (s *accountService) CreateAccount(ctx context.Context, req dto.CreateAccountRequest) error {
	if err := validator.ValidateStruct(req); err != nil {
		parsedErr := validator.ParseValidationErrors(err)
		log.Warn("Validation failed: ", parsedErr)
		return fmt.Errorf("%s", parsedErr)
	}
	return s.repo.Create(ctx, &entities.Account{
		DocumentNumber: req.DocumentNumber,
	})
}

func (s *accountService) GetAccountByID(ctx context.Context, req dto.GetAccountByIDRequest) (*dto.GetAccountByIDResponse, error) {
	if err := validator.ValidateStruct(req); err != nil {
		parsedErr := validator.ParseValidationErrors(err)
		log.Warn("Validation failed: ", parsedErr)
		return nil, fmt.Errorf("%s", parsedErr)
	}

	account, err := s.repo.GetByID(ctx, req.AccountID)

	if err != nil {
		return nil, err
	}

	return &dto.GetAccountByIDResponse{
		AccountId:      account.ID,
		DocumentNumber: account.DocumentNumber,
	}, nil
}
