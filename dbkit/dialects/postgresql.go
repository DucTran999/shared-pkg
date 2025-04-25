package dialects

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	SSLEnable  = "enable"
	SSLDisable = "disable"
)

type postgresConfig struct {
	base    Config
	SSLMode string
}

type postgresDialect struct {
	config postgresConfig
	dialect
}

func NewPostgres(cfg Config) (Dialect, error) {
	dialect := &postgresDialect{}

	if err := dialect.setupConfig(cfg); err != nil {
		return nil, err
	}

	return dialect.open()
}

func (c *postgresDialect) setupConfig(config Config) error {
	if err := validateConfig(config); err != nil {
		return err
	}

	c.config = postgresConfig{
		base:    config,
		SSLMode: SSLDisable,
	}

	if config.SSL {
		c.config.SSLMode = SSLEnable
	}

	if config.Timezone == "" {
		c.config.base.Timezone = "UTC"
	}

	return nil
}

func (c *postgresDialect) open() (Dialect, error) {
	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		c.config.base.Host, c.config.base.Username, c.config.base.Password, c.config.base.Database,
		c.config.base.Port, c.config.SSLMode, c.config.base.Timezone,
	)

	gormConfig := &gorm.Config{
		Logger: c.logger(c.config.base.Logging),
	}

	c.db, err = gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, err
	}

	return c, nil
}
