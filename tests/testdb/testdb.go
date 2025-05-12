package testdb

import (
	"pulseledger/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitTestDB(t *testing.T) *gorm.DB {
	database, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	assert.NoError(t, err)

	err = database.AutoMigrate(&entities.Account{}, &entities.Transaction{}, &entities.OperationType{})
	assert.NoError(t, err)
	return database
}
