package dbkit_test

import (
	"testing"

	"github.com/DucTran999/shared-pkg/dbkit"
	"github.com/DucTran999/shared-pkg/dbkit/dialects"
	"github.com/stretchr/testify/require"
)

func Test_PostgresConnect(t *testing.T) {
	conn, err := dbkit.NewDBDialect(
		dialects.PostgresDriver,
		dialects.Config{
			Host:     "localhost",
			Port:     5432,
			Username: "test",
			Password: "test",
			Database: "atlana_shop",
		},
	)
	require.NoError(t, err, "failed to create DB connector")
	require.NotNil(t, conn, "expected a valid DB connection")

	err = conn.Ping()
	require.NoError(t, err, "failed to open DB connection")

	defer func() {
		err := conn.Close()
		require.NoError(t, err, "failed to close DB connection")
	}()
}
