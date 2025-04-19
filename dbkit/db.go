package dbkit

import (
	"fmt"

	"gorm.io/gorm"
)

type DBConnector interface {
	DB() *gorm.DB
	Ping() error
	Open() (DBConnector, error)
	Close() error
}

func NewDBConnector(driver DBDriver, host string, options ...Option) (DBConnector, error) {
	switch driver {
	case PostgresDriver:
		return newPostgresConnector(host, options...)
	default:
		return nil, fmt.Errorf("unsupported driver: %d", driver)
	}
}
