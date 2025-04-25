package dialects

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"time"

	std_ck "github.com/ClickHouse/clickhouse-go/v2"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

type clickhouseDialect struct {
	cfg Config
	dialect
}

func NewClickhouseDialect(cfg Config) (Dialect, error) {
	if err := validateConfig(cfg); err != nil {
		return nil, err
	}

	dialect := &clickhouseDialect{
		cfg: cfg,
	}

	return dialect.open()
}

func (c *clickhouseDialect) open() (Dialect, error) {
	sqlDB := c.openClickhouseSQL()

	gormCfg := &gorm.Config{
		Logger: c.logger(c.cfg.Logging),
	}

	db, err := gorm.Open(
		clickhouse.New(clickhouse.Config{Conn: sqlDB}),
		gormCfg,
	)
	if err != nil {
		return nil, err
	}

	c.db = db
	return c, nil
}

func (c *clickhouseDialect) openClickhouseSQL() *sql.DB {
	addr := fmt.Sprintf("%s:%d", c.cfg.Host, c.cfg.Port)

	opts := &std_ck.Options{
		Addr: []string{addr},
		Auth: std_ck.Auth{
			Database: c.cfg.Database,
			Username: c.cfg.Username,
			Password: c.cfg.Password,
		},
		Settings: std_ck.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 5 * time.Second,
		Compression: &std_ck.Compression{
			Method: std_ck.CompressionLZ4,
		},
	}

	if c.cfg.SSL {
		opts.TLS = &tls.Config{InsecureSkipVerify: true}
	}

	return std_ck.OpenDB(opts)
}
