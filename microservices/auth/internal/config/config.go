package config

var conf config

type config struct {
	ServerAddr string `ENV:"SERVER_ADDR"`
}

func ServerAddr() string {
	return conf.ServerAddr
}
