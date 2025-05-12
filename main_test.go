package main

import (
	"net/http"
	"net/http/httptest"
	"pulseledger/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestConfig struct {
	DBPath           string
	ResetEnabled     bool
	Port             string
	AuditWorkerCount int
}

func (c *TestConfig) GetDBPath() string        { return c.DBPath }
func (c *TestConfig) IsDbResetEnabled() bool   { return c.ResetEnabled }
func (c *TestConfig) GetPort() string          { return c.Port }
func (c *TestConfig) GetAuditWorkerCount() int { return c.AuditWorkerCount }

func TestMainInitialization(t *testing.T) {
	config.SetConfig(&TestConfig{
		DBPath:           "file::memory:?cache=shared",
		ResetEnabled:     false,
		Port:             "3000",
		AuditWorkerCount: 1,
	})

	app := InitApp()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/accounts/1", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.True(t, resp.StatusCode == http.StatusNotFound || resp.StatusCode == http.StatusBadRequest, "Expected a valid response for /api/v1/accounts/:id")
}
