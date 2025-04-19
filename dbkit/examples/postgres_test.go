package dbkit_test

import (
	"testing"

	"github.com/DucTran999/shared-pkg/dbkit"
	"github.com/stretchr/testify/require"
)

func Test_PostgresConnect(t *testing.T) {
	conn, err := dbkit.NewDBConnector(
		dbkit.PostgresDriver,
		"localhost",
		dbkit.WithUsername("test"),
		dbkit.WithPassword("test"),
		dbkit.WithDatabase("atlana_shop"),
	)
	require.NoError(t, err, "failed to create DB connector")

	db, err := conn.Open()
	require.NoError(t, err, "failed to open DB connection")
	require.NotNil(t, db, "expected a valid DB connection")

	defer func() {
		err := db.Close()
		require.NoError(t, err, "failed to close DB connection")
	}()
}
