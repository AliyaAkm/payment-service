package epayment

import (
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Client struct {
	logger       *zap.Logger
	httpClient   *http.Client
	authHost     string
	host         string
	TerminalID   string //  идентификатор магазина
	ClientID     string //  идентификатор клиента
	ClientSecret string // секретный ключ
	SecretHash   string
}

func NewClient(logger *zap.Logger, duration time.Duration, authHost string, host string, TerminalID string, ClientID string, ClientSecret string, SecretHash string) *Client {
	return &Client{
		logger:       logger,
		httpClient:   &http.Client{Timeout: duration},
		authHost:     authHost,
		host:         host,
		TerminalID:   TerminalID,
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		SecretHash:   SecretHash,
	}
}
