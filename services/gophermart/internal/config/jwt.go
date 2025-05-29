package config

type JwtConfig struct {
	SecretKey     string
	SigningMethod string
}

func (c JwtConfig) SigningKey() []byte {
	return []byte(c.SecretKey)
}
