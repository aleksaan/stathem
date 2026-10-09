package config

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type TConfig struct {
	DbConnectionString      string        `yaml:"DB_CONN_STRING"`
	DbSslMode               string        `yaml:"DB_SSL_MODE"`
	DbSchema                string        `yaml:"DB_SCHEMA"`
	DbRecreateTables        bool          `yaml:"DB_RECREATE_TABLES"`
	DbQueryTimeout          time.Duration `yaml:"DB_QUERY_TIMEOUT"`
	AppPort                 string        `yaml:"APP_PORT"`
	AppVersion              string
	AppLimitBodySizeInBytes int64 `yaml:"APP_LIMIT_BODY_SIZE_IN_BYTES"`
}

var Config TConfig

var appVersion = "3.0.0"

func LoadConfig(path string) {
	slog.Info("Configuration loading")
	configPath := path
	if configPath == "" {
		slog.Error("Config path hasn't been found. You must set <CONFIG_PATH> environment variable")
		os.Exit(1)
	}
	slog.Info(fmt.Sprintf("--- Reading config file: %s", configPath))
	f, err := os.Open(configPath)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&Config)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	Config.AppVersion = appVersion
	slog.Info("Configuration loading...Done")
	slog.Info(fmt.Sprintf("Application version is %s", Config.AppVersion))
}
