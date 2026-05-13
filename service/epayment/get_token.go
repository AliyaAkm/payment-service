package epayment

import (
	"bytes"
	"context"
	"encoding/json"
	"go.uber.org/zap"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type GetTokenRequest struct {
	GrantType    string // указывает использование клиентской аутентификации
	Scope        string // область доступа
	ClientID     string
	ClientSecret string
	InvoiceID    string // Номер заказа
	SecretHash   string
	Amount       int
	Currency     string
	Terminal     string // Идентификатор терминала
}

type GetTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func (c *Client) GetToken(ctx context.Context, request GetTokenRequest) (GetTokenResponse, error) {
	param := url.Values{}
	param.Set("terminal", c.TerminalID)
	param.Set("client_id", c.ClientID)
	param.Set("client_secret", c.ClientSecret)
	param.Set("invoiceID", request.InvoiceID)
	param.Set("grant_type", "client_credentials")
	param.Set("scope", "webapi usermanagement email_send verification statement statistics payment")
	param.Set("secret_hash", c.SecretHash)
	if request.Currency != "" {
		param.Set("currency", request.Currency)
	}
	if request.Amount != 0 {
		param.Set("amount", strconv.Itoa(request.Amount))
	}
	c.logger.Info(
		"get token request",
		zap.String("terminal", c.TerminalID),
		zap.String("client_id", c.ClientID),
		zap.String("client_secret", c.ClientSecret),
		zap.String("invoiceID", request.InvoiceID),
		zap.String("secret_hash", c.SecretHash),
		zap.String("currency", request.Currency),
		zap.Int("amount", request.Amount),
	)
	var payload = bytes.NewBufferString(param.Encode())
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.authHost+"/oauth2/token", payload)
	if err != nil {
		return GetTokenResponse{}, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	response, err := c.httpClient.Do(req)
	if err != nil {
		return GetTokenResponse{}, err
	}
	defer response.Body.Close()

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return GetTokenResponse{}, err
	}
	c.logger.Info("epay get token response", zap.String("body", string(respBytes)), zap.Int("status", response.StatusCode))

	var result GetTokenResponse
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return GetTokenResponse{}, err
	}
	return result, nil
}
