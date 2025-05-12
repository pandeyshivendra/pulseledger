package repositories

import (
	"context"
	"pulseledger/entities"

	"gorm.io/gorm"
)

type AccountRepository interface {
	Create(ctx context.Context, account *entities.Account) error
	GetByID(ctx context.Context, id uint64) (*entities.Account, error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db}
}

func (r *accountRepository) Create(ctx context.Context, account *entities.Account) error {
	return r.db.WithContext(ctx).Create(account).Error
}

func (r *accountRepository) GetByID(ctx context.Context, id uint64) (*entities.Account, error) {
	var account entities.Account
	if err := r.db.WithContext(ctx).Take(&account, id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}
