package dbkit_test

import (
	"testing"

	"github.com/DucTran999/shared-pkg/dbkit"
	"github.com/DucTran999/shared-pkg/dbkit/dialects"
	"github.com/stretchr/testify/require"
)

func Test_ClickhouseConnect(t *testing.T) {
	conn, err := dbkit.NewDBDialect(dialects.ClickhouseDriver, dialects.Config{
		Host:     "localhost",
		Port:     9000,
		Username: "test",
		Password: "test",
		Database: "test_db",
		Logging:  true,
	})
	require.NoError(t, err, "failed to create DB connector")
	require.NotNil(t, conn, "expected a valid DB connection")

	err = conn.Ping()
	require.NoError(t, err, "failed to open DB connection")

	defer func() {
		err := conn.Close()
		require.NoError(t, err, "failed to close DB connection")
	}()
}
