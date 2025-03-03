package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env       string        `yaml:"environment" env-required:"true"`
	TokenTTL  time.Duration `yaml:"token_ttl"`
	SecretKey string        `env:"SECRET_KEY" env-required:"true"`
	Server    HTTPServer    `yaml:"http_server"`
}

type HTTPServer struct {
	Port        string        `yaml:"port"`
	Host        string        `yaml:"host"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

func MustLoad() *Config {
	_ = godotenv.Load()

	configPath := os.Getenv("CONFIG_PATH")
	if _, err := os.Stat(configPath); err != nil {
		panic(err)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic(err)
	}

	return &cfg
}
