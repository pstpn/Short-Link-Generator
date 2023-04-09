package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config - Структура, описывающая конфигурацию проекта
	Config struct {
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
	}

	// HTTP - Структура, описывающая конфигурацию сервера
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// Log - Структура, описывающая конфигурацию logger
	Log struct {
		Level string `env-required:"true" yaml:"log_level"   env:"LOG_LEVEL"`
	}
)

// NewConfig - Функция, создающая новый конфиг
func NewConfig() (*Config, error) {

	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("reading error: %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
