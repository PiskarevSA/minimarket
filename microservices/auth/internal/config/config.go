package config

import "fmt"

var conf config

type config struct {
	LogLevel           string `ENV:"LOG_LEVEL"`
	ServerAddr         string `ENV:"SERVER_ADDR"`
	PostgreSqlAddr     string
	PostgreSqlUser     string
	PostgreSqlPassword string
	PostgreSqlDb       string
	PostgreSqlSslMode  bool
	JwtSignKey         string
}

func Config() *config {
	return &conf
}

func LogLevel() string {
	return conf.LogLevel
}

func ServerAddr() string {
	return conf.ServerAddr
}

func JwtSignKey() string {
	return conf.JwtSignKey
}

func PostgreSqlConnUrl() string {
	connUrl := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		conf.PostgreSqlUser,
		conf.PostgreSqlPassword,
		conf.PostgreSqlAddr,
		conf.PostgreSqlDb,
	)

	if conf.PostgreSqlSslMode {
		connUrl += "?sslmode=disable"
	}

	return connUrl
}
