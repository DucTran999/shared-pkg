package dbkit

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
