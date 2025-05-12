package repositories

import (
	"context"
	"testing"

	"pulseledger/entities"
	"pulseledger/tests/testdb"

	"github.com/stretchr/testify/assert"
)

func TestAccountRepository_ShouldCreateAccount(t *testing.T) {
	repo := NewAccountRepository(testdb.InitTestDB(t))
	account := &entities.Account{DocumentNumber: 12345678900}

	err := repo.Create(context.Background(), account)

	assert.NoError(t, err)
	assert.NotZero(t, account.ID)
}

func TestAccountRepository_ShouldGetAccountByID(t *testing.T) {
	db := testdb.InitTestDB(t)
	repo := NewAccountRepository(db)

	seed := &entities.Account{DocumentNumber: 98765432100}
	_ = repo.Create(context.Background(), seed)

	got, err := repo.GetByID(context.Background(), seed.ID)

	assert.NoError(t, err)
	assert.Equal(t, seed.DocumentNumber, got.DocumentNumber)
}
