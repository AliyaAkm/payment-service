package epayment

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type PaymentRequest struct {
	OrderID     uuid.UUID `json:"order_id"`
	Amount      int       `json:"amount"`
	Currency    string    `json:"currency"`
	Name        string    `json:"name"`
	Cryptogram  string    `json:"cryptogram"`
	InvoiceID   string    `json:"invoiceId"`
	Description string    `json:"description"`
	CardSave    bool      `json:"cardSave"`
	PostLink    string    `json:"postLink"`
	Terminal    string    `json:"terminal"`
}

type PaymentResponse struct {
	ID                string  `json:"id"`
	Amount            int     `json:"amount"`
	AmountBonus       int     `json:"amountBonus"`
	Currency          string  `json:"currency"`
	Description       string  `json:"description"`
	InvoiceID         string  `json:"invoiceID"`
	Language          string  `json:"language"`
	Reference         string  `json:"reference"`
	IntReference      string  `json:"intReference"`
	Secure3D          any     `json:"secure3D"`
	CardID            string  `json:"cardID"`
	Fee               int     `json:"fee"`
	ApprovalCode      string  `json:"approvalCode"`
	Code              int     `json:"code"`
	Status            string  `json:"status"`
	IP                string  `json:"ip"`
	IPCity            string  `json:"ipCity"`
	IPCountry         string  `json:"ipCountry"`
	IPLatitude        float64 `json:"ipLatitude"`
	IPLongitude       float64 `json:"ipLongitude"`
	IssuerBankCountry string  `json:"issuerBankCountry"`
	IsCredit          bool    `json:"isCredit"`
	Issuer            string  `json:"issuer"`
}

func (c *Client) CreatePayment(ctx context.Context, request PaymentRequest) (PaymentResponse, error) {
	request.Description = "Course Payment"
	request.PostLink = "https://testmerchant/order/1123"
	request.Terminal = c.TerminalID

	// auth get token
	tokenResp, err := c.GetToken(ctx, GetTokenRequest{
		Currency:  request.Currency,
		Amount:    request.Amount,
		InvoiceID: request.InvoiceID,
	})
	if err != nil {
		return PaymentResponse{}, fmt.Errorf("auth resp: %w", err)
	}

	c.logger.Info(
		"epay payment request",
		zap.String("host", c.host),
		zap.String("invoiceId", request.InvoiceID),
		zap.Int("amount", request.Amount),
		zap.String("currency", request.Currency),
		zap.String("name", request.Name),
		zap.String("description", request.Description),
		zap.String("postLink", request.PostLink),
		zap.Bool("cardSave", request.CardSave),
		zap.Int("cryptogram_length", len(request.Cryptogram)),
		zap.Int("access_token_length", len(tokenResp.AccessToken)),
	)

	data, err := json.Marshal(request)
	if err != nil {
		return PaymentResponse{}, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.host+"/payment/cryptopay", bytes.NewReader(data))
	if err != nil {
		return PaymentResponse{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+tokenResp.AccessToken)

	response, err := c.httpClient.Do(req)
	if err != nil {
		return PaymentResponse{}, err
	}
	defer response.Body.Close()

	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return PaymentResponse{}, err
	}
	c.logger.Info("payment response", zap.String("body", string(respBytes)), zap.Int("status", response.StatusCode))
	var result PaymentResponse
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return PaymentResponse{}, err
	}
	return result, nil
}
