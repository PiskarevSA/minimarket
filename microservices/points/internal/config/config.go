package config

import (
	"fmt"
)

var conf config

type config struct {
	ServiceName        string
	LogLevel           string
	ServerAddr         string
	PostgreSqlAddr     string
	PostgreSqlUser     string
	PostgreSqlPassword string
	PostgreSqlDb       string
	PostgreSqlSslMode  bool
}

func Config() *config {
	return &conf
}

func ServiceName() string {
	return conf.ServiceName
}

func LogLevel() string {
	return conf.LogLevel
}

func ServerAddr() string {
	return conf.ServerAddr
}

func PostgreSqlConnUrl() string {
	connUrl := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		conf.PostgreSqlUser,
		conf.PostgreSqlPassword,
		conf.PostgreSqlAddr,
		conf.PostgreSqlDb,
	)

	if !conf.PostgreSqlSslMode {
		connUrl += "?sslmode=disable"
	}

	return connUrl
}
