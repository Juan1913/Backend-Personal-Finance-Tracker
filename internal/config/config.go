package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
}

type AppConfig struct {
	DB   DBConfig
	Port string
}

func LoadConfig() (*AppConfig, error) {
	_ = godotenv.Load()

	missing := []string{}
	get := func(key string) string {
		v := os.Getenv(key)
		if v == "" {
			missing = append(missing, key)
		}
		return v
	}

	dbConfig := DBConfig{
		Host:     get("DB_HOST"),
		User:     get("DB_USER"),
		Password: get("DB_PASSWORD"),
		Name:     get("DB_NAME"),
		Port:     get("DB_PORT"),
	}

	port := get("PORT")
	if port == "" {
		port = "8080"
	}

	if len(missing) > 0 {
		return nil, fmt.Errorf("faltan variables de entorno: %v", missing)
	}

	return &AppConfig{
		DB:   dbConfig,
		Port: port,
	}, nil
}

func (c *DBConfig) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.Host, c.User, c.Password, c.Name, c.Port)
}
