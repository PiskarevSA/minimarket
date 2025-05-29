package config

var Config config

type config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Jwt      JwtConfig
	Accrual  AccrualConfig
}
