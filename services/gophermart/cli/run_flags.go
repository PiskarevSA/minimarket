package cli

import (
	"time"

	"github.com/urfave/cli/v3"

	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/config"
)

var runFlags = []cli.Flag{
	&cli.StringFlag{
		Name:        "server.addr",
		Usage:       "Адрес HTTP сервера",
		Value:       "127.0.0.1:8616",
		Sources:     cli.EnvVars("RUN_ADDRESS"),
		Destination: &config.Config.Server.Addr,
		Aliases:     []string{"a"},
	},

	&cli.DurationFlag{
		Name:        "server.read-timeout",
		Usage:       "Таймаут чтения",
		Value:       5 * time.Second,
		Sources:     cli.EnvVars("SERVER_READ_TIMEOUT"),
		Destination: &config.Config.Server.ReadTimeout,
	},

	&cli.DurationFlag{
		Name:        "server.write-timeout",
		Usage:       "Таймаут записи",
		Value:       10 * time.Second,
		Sources:     cli.EnvVars("SERVER_WRITE_TIMEOUT"),
		Destination: &config.Config.Server.WriteTimeout,
	},

	&cli.DurationFlag{
		Name:        "server.idle-timeout",
		Usage:       "Максимальное время жизни соединения с сервером",
		Value:       120 * time.Second,
		Sources:     cli.EnvVars("SERVER_IDLE_TIMEOUT"),
		Destination: &config.Config.Server.IdleTimeout,
	},

	&cli.StringFlag{
		Name:        "postgres.uri",
		Usage:       "Унифицированный идентификатор PostgreSQL (вместо отдельных настроек)",
		Value:       "",
		Sources:     cli.EnvVars("DATABASE_URI"),
		Destination: &config.Config.Database.URI,
		Aliases:     []string{"d"},
	},

	&cli.StringFlag{
		Name:        "postgres.addr",
		Usage:       "Адрес PostgreSQL",
		Value:       "127.0.0.1:5432",
		Sources:     cli.EnvVars("POSTGRES_ADDR"),
		Destination: &config.Config.Database.Addr,
	},

	&cli.StringFlag{
		Name:        "postgres.user",
		Usage:       "Пользователь для подключения",
		Value:       "user",
		Sources:     cli.EnvVars("POSTGRES_USER"),
		Destination: &config.Config.Database.User,
	},

	&cli.StringFlag{
		Name:        "postgres.password",
		Usage:       "Пароль для подключения",
		Value:       "password",
		Sources:     cli.EnvVars("POSTGRES_PASSWORD"),
		Destination: &config.Config.Database.Password,
	},

	&cli.StringFlag{
		Name:        "postgres.db",
		Usage:       "База данных",
		Value:       "postgres",
		Sources:     cli.EnvVars("POSTGRES_DB"),
		Destination: &config.Config.Database.DB,
	},

	&cli.BoolFlag{
		Name:        "postgres.sslmode",
		Usage:       "Использовать SSL для подключения",
		Value:       false,
		Sources:     cli.EnvVars("POSTGRES_SSLMODE"),
		Destination: &config.Config.Database.SslMode,
	},

	&cli.StringFlag{
		Name:        "jwt.signkeyfilepath",
		Usage:       "Путь к файлу с ключом для подписи по указанному алгоритму",
		Value:       "jwt.pem",
		Sources:     cli.EnvVars("JWT_SECRET_KEY_PATH"),
		Destination: &config.Config.Jwt.SecretKey,
	},

	&cli.StringFlag{
		Name:        "jwt.signing-method",
		Usage:       "Алгоритм подписи",
		Value:       "ES256",
		Sources:     cli.EnvVars("JWT_SIGNING_METHOD"),
		Destination: &config.Config.Jwt.SigningMethod,
	},

	&cli.StringFlag{
		Name:        "accrual.addr",
		Usage:       "Адрес системы расчёта начислений",
		Value:       "127.0.0.1:8617",
		Sources:     cli.EnvVars("ACCRUAL_SYSTEM_ADDRESS"),
		Destination: &config.Config.Accrual.Addr,
		Aliases:     []string{"r"},
	},
}
