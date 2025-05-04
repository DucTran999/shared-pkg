package dbkit

import (
	"fmt"

	"github.com/DucTran999/shared-pkg/dbkit/dialects"
)

func NewDBDialect(driver dialects.DBDriver, config dialects.Config) (dialects.Dialect, error) {
	switch driver {
	case dialects.PostgresDriver:
		return dialects.NewPostgres(config)
	case dialects.ClickhouseDriver:
		return dialects.NewClickhouseDialect(config)
	default:
		return nil, fmt.Errorf("unsupported driver: %d", driver)
	}
}
