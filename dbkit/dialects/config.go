package dialects

import (
	"errors"
	"time"
)

// Define error messages as constants
var (
	ErrHostRequired     = errors.New("host is required")
	ErrPortRequired     = errors.New("port is required and cannot be 0")
	ErrUsernameRequired = errors.New("username is required")
	ErrDatabaseRequired = errors.New("database name is required")
)

type Config struct {
	Host string
	Port int

	Username string
	Password string
	Database string

	Logging bool

	SSL      bool
	Timezone string

	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

func validateConfig(cfg Config) error {
	if cfg.Host == "" {
		return ErrHostRequired
	}
	if cfg.Port == 0 {
		return ErrPortRequired
	}
	if cfg.Username == "" {
		return ErrUsernameRequired
	}
	if cfg.Database == "" {
		return ErrDatabaseRequired
	}

	return nil
}
