package db

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

// TODO Изменить на одну переменную (DATABASE_URL)
type DBConfig struct {
	User     string `env:"DB_USER" env-required:"true"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
	Host     string `env:"DB_HOST" env-required:"true"`
	Port     string `env:"DB_PORT" env-required:"true"`
	Name     string `env:"DB_NAME" env-required:"true"`
	SslMode  string `env:"DB_SSLMODE" env-default:"disable"`
}

func GetConfig() string {
	var cfg DBConfig

	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		log.Fatalf("Error reading .env: %v", err)
	}

	config := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SslMode)

	return config
}
