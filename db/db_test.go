package db

import (
	"pulseledger/config"
	"pulseledger/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	DBPath           string
	ResetEnabled     bool
	Port             string
	AuditWorkerCount int
}

func (c *TestConfig) GetDBPath() string {
	return c.DBPath
}
func (c *TestConfig) IsDbResetEnabled() bool {
	return c.ResetEnabled
}
func (c *TestConfig) GetPort() string {
	return c.Port
}
func (c *TestConfig) GetAuditWorkerCount() int {
	return c.AuditWorkerCount
}

func TestInit_ShouldInitializeDatabase(t *testing.T) {
	config.SetConfig(&TestConfig{
		DBPath:           "file::memory:?cache=shared",
		ResetEnabled:     false,
		Port:             "8080",
		AuditWorkerCount: 1,
	})

	db := Init()
	assert.NotNil(t, db, "Expected initialized database to be non-nil")

	result := GetDatabase()
	assert.Equal(t, db, result, "Expected GetDatabase to return the same instance")
}

func TestInitWithConfig_ShouldResetAndMigrate(t *testing.T) {
	cfg := &TestConfig{
		DBPath:           "file::memory:?cache=shared",
		ResetEnabled:     true,
		Port:             "8080",
		AuditWorkerCount: 1,
	}
	db := InitWithConfig(cfg)

	assert.NotNil(t, db, "Expected database to be initialized with reset config")

	assert.True(t, db.Migrator().HasTable(&entities.Account{}), "Expected accounts table to be created")
	assert.True(t, db.Migrator().HasTable(&entities.Transaction{}), "Expected transactions table to be created")
	assert.True(t, db.Migrator().HasTable(&entities.OperationType{}), "Expected operation_types table to be created")
}
