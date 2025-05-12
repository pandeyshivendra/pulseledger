package repositories

import (
	"context"
	"testing"
	"time"

	"pulseledger/entities"
	"pulseledger/enums"
	"pulseledger/tests/testdb"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRepository_ShouldCreateTransaction(t *testing.T) {
	db := testdb.InitTestDB(t)

	account := &entities.Account{DocumentNumber: 11122233344}
	err := db.Create(account).Error
	assert.NoError(t, err)

	operationType := &entities.OperationType{
		ID:           uint8(enums.NormalPurchase),
		Description:  "Normal Purchase",
		IsDepricated: false,
	}
	err = db.Create(operationType).Error
	assert.NoError(t, err)

	repo := NewTransactionRepository(db)
	transaction := &entities.Transaction{
		AccountID:       account.ID,
		OperationTypeID: enums.NormalPurchase,
		Amount:          decimal.NewFromFloat(-123.45),
		EventDate:       time.Now().UTC(),
	}

	err = repo.Create(context.Background(), transaction)

	assert.NoError(t, err)
	assert.NotZero(t, transaction.ID)
}
