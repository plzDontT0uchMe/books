package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config содержит параметры конфигурации приложения

type Config struct {
	Environment    string `env:"ENVIRONMENT" env-required:"true"`
	ServiceName    string `env:"SERVICE_NAME" env-required:"true"`
	ServiceVersion string `env:"SERVICE_VERSION" env-required:"true"`
	Server
}

type Server struct {
	Protocol          string        `env:"SERVER_PROTOCOL" env-required:"true"`
	Host              string        `env:"SERVER_HOST" env-required:"true"`
	Port              int           `env:"SERVER_PORT" env-required:"true"`
	Address           string        `env:"SERVER_ADDRESS" env-required:"true"`
	RequestTimeout    time.Duration `env:"SERVER_REQUEST_TIMEOUT" env-required:"true"`
	RateLimitCount    int           `env:"SERVER_RATE_LIMIT_COUNT" env-required:"true"`
	RateLimitDuration time.Duration `env:"SERVER_RATE_LIMIT_DURATION" env-required:"true"`
}

// MustLoad загружает конфигурацию из файла, указанного либо по ключу 'config' в командной строке, либо
// в переменной окружения 'CONFIG_PATH'. Если путь к файлу не указан, то используются только переменные окружения.
// Параметры конфигурации из файла замещаются соответствующими параметрами из переменных окружения при их наличии.
func MustLoad() (*Config, error) {
	var configPath = flag.String("config", "", "путь к файлу конфигурации")
	var cfg Config

	flag.Parse()

	if *configPath == "" {
		*configPath = os.Getenv("CONFIG_PATH")
	}

	// Временно для быстрого запуска
	if *configPath == "" {
		*configPath = "./configs/local.env"
	}

	if *configPath == "" {
		if err := cleanenv.ReadEnv(&cfg); err != nil {
			return nil, err
		}
	} else {
		if _, err := os.Stat(*configPath); os.IsNotExist(err) {
			return nil, err
		}

		if err := cleanenv.ReadConfig(*configPath, &cfg); err != nil {
			return nil, fmt.Errorf("cannot read config: %s", err)
		}
	}

	return &cfg, nil
}
