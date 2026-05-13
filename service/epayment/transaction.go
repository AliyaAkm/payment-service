package epayment

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io"
	"net/http"
	"time"
)

type GetTransactionResponse struct {
	ResultCode    string       `json:"resultCode"`
	ResultMessage string       `json:"resultMessage"`
	Transaction   *Transaction `json:"transaction"`
}

type Transaction struct {
	ID          string    `json:"id"`
	CreatedDate time.Time `json:"createdDate"`
	InvoiceID   string    `json:"invoiceID"`

	Amount       float64 `json:"amount"`
	AmountBonus  float64 `json:"amountBonus"`
	PayoutAmount float64 `json:"payoutAmount"`
	OrgAmount    float64 `json:"orgAmount"`

	ApprovalCode string `json:"approvalCode"`
	Data         string `json:"data"`

	Currency    string `json:"currency"`
	Terminal    string `json:"terminal"`
	TerminalID  string `json:"terminalID"`
	AccountID   string `json:"accountID"`
	Description string `json:"description"`
	Language    string `json:"language"`

	CardMask string `json:"cardMask"`
	CardType string `json:"cardType"`
	Issuer   string `json:"issuer"`

	Reference    string `json:"reference"`
	Reason       string `json:"reason"`
	ReasonCode   string `json:"reasonCode"`
	IntReference string `json:"intReference"`

	Secure        bool   `json:"secure"`
	SecureDetails string `json:"secureDetails"`

	StatusID   string `json:"statusID"`
	StatusName string `json:"statusName"`

	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`

	CardID string `json:"cardID"`
	XLSRRN string `json:"xlsRRN"`

	IP          string  `json:"ip"`
	IPCountry   string  `json:"ipCountry"`
	IPCity      string  `json:"ipCity"`
	IPRegion    string  `json:"ipRegion"`
	IPDistrict  string  `json:"ipDistrict"`
	IPLatitude  float64 `json:"ipLatitude"`
	IPLongitude float64 `json:"ipLongitude"`
}

func (c *Client) GetStatusTransaction(ctx context.Context, invoiceID string) (GetTransactionResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.host+"/check-status/payment/transaction/"+invoiceID, nil)
	if err != nil {
		return GetTransactionResponse{}, err
	}

	tokenResp, err := c.GetToken(ctx, GetTokenRequest{

		InvoiceID: invoiceID,
	})

	req.Header.Add("Authorization", "Bearer "+tokenResp.AccessToken)

	response, err := c.httpClient.Do(req)
	if err != nil {
		return GetTransactionResponse{}, err
	}
	defer response.Body.Close()
	respBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return GetTransactionResponse{}, err
	}
	c.logger.Info("transaction get response", zap.String("body", string(respBytes)), zap.Int("status", response.StatusCode))
	if response.StatusCode != http.StatusOK {
		// to do: обработка ошибки через unmarshal
		return GetTransactionResponse{}, fmt.Errorf("something goes wrong from API: %w", err)
	}
	var result GetTransactionResponse
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		return GetTransactionResponse{}, err
	}
	return result, nil

}
