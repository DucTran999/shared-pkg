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
	var err error

	gormCfg := &gorm.Config{
		Logger: c.logger(c.cfg.Logging),
	}

	ch := clickhouse.New(clickhouse.Config{
		Conn: c.openClickhouseSQL(),
	})
	if c.db, err = gorm.Open(ch, gormCfg); err != nil {
		return nil, err
	}

	if err = c.dialect.configPool(c.cfg); err != nil {
		return nil, err
	}

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
		opts.TLS = &tls.Config{
			MinVersion:         tls.VersionTLS12,
			InsecureSkipVerify: true,
		}
	}

	return std_ck.OpenDB(opts)
}
