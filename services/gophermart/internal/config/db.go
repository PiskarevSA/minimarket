package config

import (
	"fmt"
)

type DatabaseConfig struct {
	Addr     string
	User     string
	Password string
	Db       string
	SslMode  bool
}

func (c DatabaseConfig) ConnUrl() string {
	connUrl := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		c.User,
		c.Password,
		c.Addr,
		c.Db,
	)

	if c.SslMode {
		connUrl += "?sslmode=disable"
	}

	return connUrl
}
