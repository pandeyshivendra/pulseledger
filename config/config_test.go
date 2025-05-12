package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestLoad_WithMissingConfigFile_ShouldFallbackToEnv(t *testing.T) {
	os.Setenv("SERVER_PORT", "8081")
	os.Setenv("DATABASE_PATH", ":memory:")
	os.Setenv("DATABASE_RESET_ON_START", "true")
	os.Setenv("AUDIT_WORKER_COUNT", "5")

	viper.Reset()
	Load()

	cfg := GetConfig()
	assert.Equal(t, "8081", cfg.GetPort())
	assert.Equal(t, ":memory:", cfg.GetDBPath())
	assert.True(t, cfg.IsDbResetEnabled())
	assert.Equal(t, 5, cfg.GetAuditWorkerCount())
}

func TestMustGetString_MissingKey_ShouldPanic(t *testing.T) {
	viper.Reset()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for missing string key")
		}
	}()
	mustGetString("missing.string.key")
}

func TestMustGetInt_MissingKey_ShouldPanic(t *testing.T) {
	viper.Reset()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for missing int key")
		}
	}()
	mustGetInt("missing.int.key")
}

func TestMustGetBool_MissingKey_ShouldPanic(t *testing.T) {
	viper.Reset()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic for missing bool key")
		}
	}()
	mustGetBool("missing.bool.key")
}
