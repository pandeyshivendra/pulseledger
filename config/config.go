package config

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config interface {
	GetPort() string
	GetDBPath() string
	IsDbResetEnabled() bool
	GetAuditWorkerCount() int
}

type viperConfig struct{}

var defaultConfig Config = &viperConfig{}

func SetConfig(c Config) {
	defaultConfig = c
}

func GetConfig() Config {
	return defaultConfig
}

func Load() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Warnf("No config file found. Falling back to environment variables. Error: %v", err)
	}
}

func (v *viperConfig) GetPort() string {
	return mustGetString("server.port")
}

func (v *viperConfig) GetDBPath() string {
	return mustGetString("database.path")
}

func (v *viperConfig) IsDbResetEnabled() bool {
	return mustGetBool("database.reset_on_start")
}

func (v *viperConfig) GetAuditWorkerCount() int {
	return mustGetInt("audit.worker.count")
}

func mustHaveKey(key string) {
	if !viper.IsSet(key) {
		log.Errorf("Missing key '%s'", key)
		panic("Missing key '" + key + "'")
	}
}

func mustGetString(key string) string {
	mustHaveKey(key)
	return viper.GetString(key)
}

func mustGetInt(key string) int {
	mustHaveKey(key)
	return viper.GetInt(key)
}

func mustGetBool(key string) bool {
	mustHaveKey(key)
	return viper.GetBool(key)
}
