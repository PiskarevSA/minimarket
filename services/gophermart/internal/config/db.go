package config

import (
	"fmt"
)

type DatabaseConfig struct {
	URI      string
	Addr     string
	User     string
	Password string
	DB       string
	SslMode  bool
}

func (c DatabaseConfig) ConnURL() string {
	if len(c.URI) > 0 {
		return c.URI
	}

	connURL := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		c.User,
		c.Password,
		c.Addr,
		c.DB,
	)

	if c.SslMode {
		connURL += "?sslmode=disable"
	}

	return connURL
}
