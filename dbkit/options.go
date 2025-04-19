package dbkit

import (
	"errors"
	"strings"
)

type SSLMode string

var Enable SSLMode = "enable"
var Disable SSLMode = "disable"

type options struct {
	host string
	port int

	username string
	password string
	database string

	ssl      SSLMode
	timezone string
}

type Option func(option *options) error

func WithPort(port int) Option {
	return func(options *options) error {
		options.port = port

		return nil
	}
}

func WithUsername(username string) Option {
	return func(options *options) error {
		username = strings.Trim(username, " ")
		if len(username) == 0 {
			return errors.New("missing username")
		}

		options.username = username
		return nil
	}
}

func WithPassword(password string) Option {
	return func(options *options) error {
		password = strings.Trim(password, " ")
		options.password = password

		return nil
	}
}

func WithDatabase(database string) Option {
	return func(options *options) error {
		database = strings.Trim(database, " ")
		if len(database) == 0 {
			return errors.New("missing database name")
		}

		options.database = database

		return nil
	}
}

func WithSsl(enable SSLMode) Option {
	return func(options *options) error {
		options.ssl = enable
		return nil
	}
}

func WithTimezone(timezone string) Option {
	return func(options *options) error {
		options.timezone = timezone
		return nil
	}
}
