package config

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Port string
	}
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Name     string
	}
}

func Load() *Config {
	readConfig("internal/config", "config", "yaml")
	mergeOptionalEnvFile(".env")
	viper.AutomaticEnv()
	return unmarshalConfig()
}

func readConfig(path, name, configType string) {
	viper.SetConfigName(name)
	viper.SetConfigType(configType)
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("failed to read config:", err)
	}
}

func mergeOptionalEnvFile(filePath string) {
	if _, err := os.Stat(filePath); err == nil {
		viper.SetConfigFile(filePath)
		viper.SetConfigType("env")
		if err := viper.MergeInConfig(); err != nil {
			log.Fatal("failed to read .env:", err)
		}
		return
	} else if !errors.Is(err, os.ErrNotExist) {
		log.Fatal("failed to stat .env:", err)
	}
}

func unmarshalConfig() *Config {
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal("failed to unmarshal config:", err)
	}

	if envUser := viper.GetString("DB_USER"); envUser != "" {
		cfg.Database.User = envUser
	}
	if envPassword := viper.GetString("DB_PASSWORD"); envPassword != "" {
		cfg.Database.Password = envPassword
	}

	return &cfg
}
