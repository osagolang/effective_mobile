package db

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

// TODO Изменить на одну переменную (DATABASE_URL)
type DBConfig struct {
	User     string `env:"DB_USER" env-default:"eff"`
	Password string `env:"DB_PASSWORD" env-default:"mob"`
	Host     string `env:"DB_HOST" env-default:"database"`
	Port     string `env:"DB_PORT" env-default:"5432"`
	Name     string `env:"DB_NAME" env-default:"eff_mob"`
	SslMode  string `env:"DB_SSLMODE" env-default:"disable"`
}

func GetConfig() string {
	var cfg DBConfig

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		log.Fatalf("Error reading .env: %v", err)
	}

	config := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SslMode)

	return config
}
