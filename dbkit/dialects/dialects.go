package dialects

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBDriver int

const (
	PostgresDriver DBDriver = iota
	MySQLDriver
	ClickhouseDriver
)

func (d DBDriver) String() string {
	switch d {
	case PostgresDriver:
		return "postgres"
	case MySQLDriver:
		return "mysql"
	case ClickhouseDriver:
		return "clickhouse"
	default:
		return "unknown"
	}
}

type Dialect interface {
	DB() *gorm.DB

	Ping() error

	Close() error
}

type dialect struct {
	db *gorm.DB
}

func (c *dialect) DB() *gorm.DB {
	return c.db
}

func (c *dialect) Ping() error {
	sqlDB, err := c.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Ping()
}

func (c *dialect) Close() error {
	if c.db == nil {
		return nil
	}

	sqlDB, err := c.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

func (c *dialect) logger(enable bool) logger.Interface {
	if enable {
		return logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             200 * time.Millisecond,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				Colorful:                  true,
			},
		)
	}

	return nil
}

func (c *dialect) configPool(config Config) error {
	if c.db == nil {
		return nil
	}

	sqlDB, err := c.db.DB()
	if err != nil {
		return err
	}

	// Default connection pool configuration
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if config.MaxIdleConns > 0 {
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	}
	if config.ConnMaxLifetime > 0 {
		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(config.ConnMaxLifetime)
	}
	if config.MaxOpenConns > 0 {
		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	}

	return nil
}
