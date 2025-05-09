package config

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var conf config

type config struct {
	LogLevel           string `ENV:"LOG_LEVEL"`
	ServerAddr         string `ENV:"SERVER_ADDR"`
	PostgreSqlAddr     string
	PostgreSqlUser     string
	PostgreSqlPassword string
	PostgreSqlDb       string
	PostgreSqlSslMode  bool
	JwtSignKeyFilePath string
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

func JwtSignKeyFilePath() string {
	return conf.JwtSignKeyFilePath
}

func JwtAlgo() jwt.SigningMethod {
	return jwt.SigningMethodES256
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
