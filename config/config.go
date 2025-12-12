package config

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}
type ServerConfig struct {
	Port string
}
type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
func Validate(c *Config) error {
	if c.Database.Password == "" {
		return errors.New("DB_PASS is empty")
	}
	if c.Database.Host == "" {
		return errors.New("DB_HOST is empty")
	}
	if c.Database.DBName == "" {
		return errors.New("DB_NAME is empty")
	}
	return nil
}

func Load() (*Config, error) {
	cfg := &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
		},
		Database: DatabaseConfig{
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "5432"),
			Username: getEnv("DB_USER", "postgres"),
			Password: getEnv("DB_PASS", "postgres"),
			DBName:   getEnv("DB_NAME", "postgres"),
		},
	}
	if err := Validate(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func (d *DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", d.Host, d.Port, d.Username, d.Password, d.DBName)
}
