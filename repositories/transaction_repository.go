package repositories

import (
	"context"
	"pulseledger/entities"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(ctx context.Context, transaction *entities.Transaction) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) Create(ctx context.Context, tx *entities.Transaction) error {
	return r.db.WithContext(ctx).Create(tx).Error
}
