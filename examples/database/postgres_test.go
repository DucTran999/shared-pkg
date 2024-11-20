package gorm_test

import (
	"testing"

	gorm "github.com/DucTran999/shared-pkg/database"
	"github.com/stretchr/testify/assert"
)

func TestPostgresConnect(t *testing.T) {
	conf := gorm.DBConfig{
		Driver:                gorm.PostgresDriver,
		Env:                   "test",
		Host:                  "localhost",
		Port:                  5432,
		Username:              "test",
		Password:              "test",
		Database:              "atlana_shop",
		SslMode:               "disable",
		Timezone:              "Asia/Ho_Chi_Minh",
		MaxOpenConnections:    500,
		MaxIdleConnections:    5,
		MaxConnectionIdleTime: 10,
	}
	conn := gorm.NewDBConnector(conf)

	pg, err := conn.Connect()

	assert.Nil(t, err)
	pg.Stop()
}
