package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}
type ServerConfig struct {
	Port              string
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	ReadTimeout       time.Duration
	IdleTimeout       time.Duration
	ShutdownTimeout   time.Duration
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
func getEnvDuration(key string, defaultValue time.Duration) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		if d, err := time.ParseDuration(value); err == nil {
			return d
		}
	}
	return defaultValue
}

func Validate(c *Config) error {
	return nil
}

func Load() (*Config, error) {
	cfg := &Config{
		Server: ServerConfig{
			Port:              getEnv("SERVER_PORT", "8080"),
			ReadHeaderTimeout: getEnvDuration("SERVER_READ_HEADER_TIMEOUT", 5*time.Second),
			WriteTimeout:      getEnvDuration("SERVER_WRITE_TIMEOUT", 10*time.Second),
			ReadTimeout:       getEnvDuration("SERVER_READ_TIMEOUT", 5*time.Second),
			IdleTimeout:       getEnvDuration("SERVER_IDLE_TIMEOUT", 30*time.Second),
			ShutdownTimeout:   getEnvDuration("SERVER_SHUTDOWN_TIMEOUT", 10*time.Second),
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
