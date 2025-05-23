package config

type JwtConfig struct {
	SecretKey string
}

func (c JwtConfig) SigningKey() []byte {
	return []byte(c.SecretKey)
}
