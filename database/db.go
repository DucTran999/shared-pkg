package gorm

import (
	"gorm.io/gorm"
)

type IDBConnector interface {
	GetConn() *gorm.DB
	Connect() (IDBConnector, error)
	Stop() error
}

type DBConfig struct {
	Driver                string
	Env                   string
	Host                  string
	Port                  int
	Username              string
	Password              string
	Database              string
	SslMode               string
	Timezone              string
	MaxOpenConnections    int
	MaxIdleConnections    int
	MaxConnectionIdleTime int
}

const (
	PostgresDriver = "postgres"
)

func NewDBConnector(conf DBConfig) IDBConnector {
	switch conf.Driver {
	case PostgresDriver:
		return newPostgresConnector(conf)
	default:
		return nil
	}
}
