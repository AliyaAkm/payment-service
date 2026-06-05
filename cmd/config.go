package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"github.com/caarlos0/env/v11"
)

type HTTPConfig struct {
	ReadHeaderTimeout time.Duration `env:"READ_HEADER_TIMEOUT" envDefault:"5s"`
	ShutdownTimeout   time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"5s"`
}

type DBConfig struct {
	Host              string        `env:"HOST"`
	Port              int           `env:"PORT"`
	User              string        `env:"USER"`
	Password          string        `env:"PASSWORD"`
	DBName            string        `env:"NAME"`
	SSLMode           string        `env:"SSLMODE"`
	MaxConns          int32         `env:"MAX_CONNS"`
	MinConns          int32         `env:"MIN_CONNS"`
	MaxConnLifetime   time.Duration `env:"MAX_CONN_LIFETIME"`
	HealthCheckPeriod time.Duration `env:"HEALTH_CHECK_PERIOD"`
}

type EpayConfig struct {
	Duration     time.Duration `env:"DURATION"`
	AuthHost     string        `env:"AUTH_HOST"`
	Host         string        `env:"HOST"`
	TerminalID   string        `env:"TERMINAL_ID"`
	ClientSecret string        `env:"CLIENT_SECRET"`
	ClientID     string        `env:"CLIENT_ID"`
	SecretHash   string        `env:"SECRET_HASH"`
	PublicKey    string        `env:"PUBLIC_KEY"`
}

type NotificationServiceConfig struct {
	URL            string        `env:"URL" envDefault:"http://localhost:8087"`
	Timeout        time.Duration `env:"TIMEOUT" envDefault:"5s"`
	InternalAPIKey string        `env:"INTERNAL_API_KEY" envDefault:""`
}

type Config struct {
	HTTPAddr     string                    `env:"HTTP_ADDR" envDefault:":8080"`
	HTTP         HTTPConfig                `envPrefix:"HTTP_"`
	DB           DBConfig                  `envPrefix:"DB_"`
	Epayment     EpayConfig                `envPrefix:"EPAYMENT_"`
	Notification NotificationServiceConfig `envPrefix:"NOTIFICATION_SERVICE_"`
}

func ReadEnv() (*Config, error) {
	cfg := new(Config)
	opts := env.Options{
		RequiredIfNoDef: true,
	}

	if err := env.ParseWithOptions(cfg, opts); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c Config) ListenAddr() string {
	port := strings.TrimSpace(os.Getenv("PORT"))
	if port == "" {
		return c.HTTPAddr
	}

	return net.JoinHostPort("", port)
}

func (db DBConfig) DatabaseURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.DBName,
		db.SSLMode,
	)
}
