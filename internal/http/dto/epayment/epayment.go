package epayment

import (
	"github.com/google/uuid"
	"time"
)

type GetTokenRequest struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

type GetTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type PaymentRequest struct {
	OrderID    uuid.UUID `json:"order_id"`
	Name       string    `json:"name"`
	Cryptogram string    `json:"сryptogram"`
	CardSave   bool      `json:"cardsave"`
}

type PaymentResponse struct {
	ID                string  `json:"id"`
	AccountID         string  `json:"accountId"`
	Amount            int     `json:"amount"`
	AmountBonus       int     `json:"amountBonus"`
	Currency          string  `json:"currency"`
	Description       string  `json:"description"`
	Email             string  `json:"email"`
	InvoiceID         string  `json:"invoiceID"`
	Language          string  `json:"language"`
	Phone             string  `json:"phone"`
	Reference         string  `json:"reference"`
	IntReference      string  `json:"intReference"`
	Secure3D          any     `json:"secure3D"`
	CardID            string  `json:"cardID"`
	Fee               int     `json:"fee"`
	ApprovalCode      string  `json:"approvalCode"`
	Code              int     `json:"code"`
	Status            string  `json:"status"`
	SecureDetails     string  `json:"secureDetails"`
	QRReference       string  `json:"qrReference"`
	IP                string  `json:"ip"`
	IPCity            string  `json:"ipCity"`
	IPCountry         string  `json:"ipCountry"`
	IPDistrict        string  `json:"ipDistrict"`
	IPLatitude        float64 `json:"ipLatitude"`
	IPLongitude       float64 `json:"ipLongitude"`
	IPRegion          string  `json:"ipRegion"`
	IssuerBankCountry string  `json:"issuerBankCountry"`
	IsCredit          bool    `json:"isCredit"`
	Issuer            string  `json:"issuer"`
}

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
