package config

import (
	"github.com/caarlos0/env/v10"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ServerPort string `env:"SERVER_PORT,required"`
	DBHost     string `env:"DB_HOST,required"`
	DBName     string `env:"DB_NAME,required"`
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBSSLMode  string `env:"DB_SSLMODE,required"`

	JWTSecret string `env:"JWT_SECRET,required"`

	EmailSender   string `env:"EMAIL_USER,required"`
	EmailPassword string `env:"EMAIL_PASS,required"`
}

func NewEnvConfig() *EnvConfig {
	if err := godotenv.Load(); err != nil {
		log.Warn(".env not found, using system environment")
	}

	cfg := &EnvConfig{}

	if err := env.Parse(cfg); err != nil {
		log.Fatalf("Unable to parse env: %v", err)
	}

	return cfg
}
