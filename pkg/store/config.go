package store

import (
	"fmt"
	"strings"

	"github.com/gofor-little/env"
)

type Config struct {
	Host     string
	Port     string
	Database string
	User     string
	Password string
}

func NewConfig() (*Config, error) {
	return &Config{
		Host:     env.Get("POSTGRES_HOST", "localhost"),
		Port:     env.Get("POSTGRES_PORT", "5432"),
		Database: env.Get("POSTGRES_DB", "postgres"),
		User:     env.Get("POSTGRES_USER", "postgres"),
		Password: env.Get("POSTGRES_PASSWORD", ""),
	}, nil
}

func (config *Config) CreateURL() string {
	var b strings.Builder

	b.WriteString("postgres://")
	b.WriteString(config.User)
	b.WriteRune(':')
	b.WriteString(config.Password)
	b.WriteRune('@')
	b.WriteString(config.Host)
	b.WriteRune(':')
	b.WriteString(config.Port)
	b.WriteRune('/')
	b.WriteString(config.Database)

	fmt.Println(b.String())

	return b.String()
}
