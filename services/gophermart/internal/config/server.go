package config

import "time"

type ServerConfig struct {
	Addr         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
}
