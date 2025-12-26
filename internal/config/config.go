package config

import "time"

type Config struct {
	DB  DBConfig
	JWT JWTConfig
}

type DBConfig struct {
	DSN string
}

type JWTConfig struct {
	AccessSecret  string
	RefreshSecret string
	AccessTTL     time.Duration
	RefreshTTL    time.Duration
}

func Load() *Config {
	return &Config{
		DB: DBConfig{
			DSN: "postgres://postgres:postgres@localhost:5432/auth?sslmode=disable",
		},
		JWT: JWTConfig{
			AccessSecret:  "access-secret",
			RefreshSecret: "refresh-secret",
			AccessTTL:     time.Minute * 15,
			RefreshTTL:    time.Hour * 24 * 7,
		},
	}
}
