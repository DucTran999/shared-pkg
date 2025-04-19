package dbkit

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresConnector struct {
	db     *gorm.DB
	config options
}

func newPostgresConnector(host string, opts ...Option) (*postgresConnector, error) {
	option := options{
		host:     host,
		port:     5432,
		ssl:      Disable,
		timezone: "UTC",
	}

	for _, opt := range opts {
		if err := opt(&option); err != nil {
			return nil, err
		}
	}

	return &postgresConnector{
		config: option,
	}, nil
}

func (c *postgresConnector) DB() *gorm.DB {
	return c.db
}

func (c *postgresConnector) Open() (DBConnector, error) {
	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		c.config.host, c.config.username, c.config.password, c.config.database,
		c.config.port, c.config.ssl, c.config.timezone,
	)

	c.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *postgresConnector) Ping() error {
	sqlDB, err := c.db.DB()
	if err != nil {
		return err
	}

	if err := sqlDB.Ping(); err != nil {
		return err
	}

	return nil
}

func (c *postgresConnector) Close() error {
	if c.db == nil {
		return nil
	}

	sqlDB, err := c.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
