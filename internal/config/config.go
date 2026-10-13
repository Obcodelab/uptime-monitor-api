// Package config reads all settings from environment variables in one place.
package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	AppEnv              string
	HTTPPort            string
	DatabaseURL         string
	JWTSecret           string
	CheckerWorkers      int
	CheckerTickSeconds  int
	CheckRetentionDays  int
	AllowPrivateTargets bool
}

// Load parses the environment, applies defaults and fails fast on bad values.
func Load() (Config, error) {
	var errs []error
	stringVar := func(key, def string) string {
		return getString(key, def)
	}
	intVar := func(key string, def int) int {
		v, err := getInt(key, def)
		if err != nil {
			errs = append(errs, err)
		}
		return v
	}
	boolVar := func(key string, def bool) bool {
		v, err := getBool(key, def)
		if err != nil {
			errs = append(errs, err)
		}
		return v
	}

	cfg := Config{
		AppEnv:              stringVar("APP_ENV", "development"),
		HTTPPort:            stringVar("HTTP_PORT", "8000"),
		DatabaseURL:         stringVar("DATABASE_URL", ""),
		JWTSecret:           stringVar("JWT_SECRET", ""),
		CheckerWorkers:      intVar("CHECKER_WORKERS", 10),
		CheckerTickSeconds:  intVar("CHECKER_TICK_SECONDS", 5),
		CheckRetentionDays:  intVar("CHECK_RETENTION_DAYS", 30),
		AllowPrivateTargets: boolVar("ALLOW_PRIVATE_TARGETS", false),
	}

	if cfg.DatabaseURL == "" {
		errs = append(errs, errors.New("DATABASE_URL is required"))
	}
	if len(cfg.JWTSecret) < 32 {
		errs = append(errs, errors.New("JWT_SECRET is required and must be at least 32 characters"))
	}
	if cfg.CheckerWorkers < 1 {
		errs = append(errs, errors.New("CHECKER_WORKERS must be at least 1"))
	}
	if cfg.CheckerTickSeconds < 1 {
		errs = append(errs, errors.New("CHECKER_TICK_SECONDS must be at least 1"))
	}
	if cfg.CheckRetentionDays < 1 {
		errs = append(errs, errors.New("CHECK_RETENTION_DAYS must be at least 1"))
	}

	if err := errors.Join(errs...); err != nil {
		return Config{}, fmt.Errorf("invalid configuration: %w", err)
	}
	return cfg, nil
}

func getString(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func getInt(key string, def int) (int, error) {
	v := os.Getenv(key)
	if v == "" {
		return def, nil
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return def, fmt.Errorf("%s must be an integer: %w", key, err)
	}
	return n, nil
}

func getBool(key string, def bool) (bool, error) {
	v := os.Getenv(key)
	if v == "" {
		return def, nil
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return def, fmt.Errorf("%s must be a boolean: %w", key, err)
	}
	return b, nil
}
