package config

var conf config

type config struct {
	LogLevel   string `ENV:"LOG_LEVEL"`
	ServerAddr string `ENV:"SERVER_ADDR"`
}

func Config() *config {
	return &conf
}

func ServerAddr() string {
	return conf.ServerAddr
}
