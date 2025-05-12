package services

import (
	"context"
	"fmt"
	"pulseledger/dto"
	"pulseledger/entities"
	"pulseledger/repositories"
	"pulseledger/validator"
	"time"

	log "github.com/sirupsen/logrus"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, req dto.CreateTransactionRequest) error
}

type transactionService struct {
	repo repositories.TransactionRepository
}

func NewTransactionService(repo repositories.TransactionRepository) TransactionService {
	return &transactionService{repo}
}

func (s *transactionService) CreateTransaction(ctx context.Context, req dto.CreateTransactionRequest) error {
	if err := validator.ValidateStruct(req); err != nil {
		parsedErr := validator.ParseValidationErrors(err)
		log.Warn("Validation failed: ", parsedErr)
		return fmt.Errorf("%s", parsedErr)
	}

	tx := entities.Transaction{
		AccountID:       req.AccountID,
		OperationTypeID: req.OperationTypeID,
		Amount:          req.Amount,
		EventDate:       time.Now().UTC(),
	}

	return s.repo.Create(ctx, &tx)
}
