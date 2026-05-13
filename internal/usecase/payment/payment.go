package payment

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"payment-service/internal/domain/order"
	"payment-service/internal/domain/orderstatus"
	domainpayment "payment-service/internal/domain/payment"
	"payment-service/internal/domain/subscription"
	"payment-service/service/epayment"
	"strconv"
	"strings"
	"time"
)

const statusPaid = "paid"
const statusCanceled = "canceled"

type CryptogramPayload struct {
	HPAN       string `json:"hpan"`
	ExpDate    string `json:"expDate"`
	CVC        string `json:"cvc"`
	TerminalID string `json:"terminalId"`
}

func (u *UseCase) CreatePayment(ctx context.Context, request epayment.PaymentRequest) (*order.Order, error) {
	order, err := u.orderRepo.GetOrderByID(ctx, request.OrderID)
	if err != nil {
		return nil, err
	}
	request.Amount = order.Amount
	request.Currency = order.Currency
	request.InvoiceID = generateInvoiceID()

	if request.Cryptogram == "" {
		publicKeyPEM := []byte(strings.ReplaceAll(u.PublicKey, `\n`, "n"))

		cryptogram, err := GenerateCryptogram(publicKeyPEM, CryptogramPayload{
			HPAN:       "377514500004820",
			ExpDate:    "0128",
			CVC:        "0198",
			TerminalID: u.TerminalID,
		})
		if err != nil {
			return nil, err
		}
		request.Cryptogram = cryptogram
	}

	_, err = u.paymentClient.CreatePayment(ctx, request)
	if err != nil {
		return nil, err
	}
	transactionInfo, err := u.paymentClient.GetStatusTransaction(ctx, request.InvoiceID)
	if err != nil {
		return nil, err
	}
	var status *orderstatus.OrderStatus
	if transactionInfo.ResultCode == "100" {
		status, err = u.orderStatusRepo.GetOrderStatusByCode(ctx, statusPaid)
		if err != nil {
			return nil, err
		}
		order.StatusID = status.ID

		subscription := subscription.Subscription{
			ID:       uuid.New(),
			UserID:   order.UserID,
			CourseID: order.CourseID,
		}
		err = u.subscriptionRepo.CreateSubscription(ctx, &subscription)
		if err != nil {
			return nil, err
		}
	} else {
		status, err = u.orderStatusRepo.GetOrderStatusByCode(ctx, statusCanceled)
		if err != nil {
			return nil, err
		}
		order.StatusID = status.ID
	}
	err = u.orderRepo.UpdateOrderStatus(ctx, request.OrderID, status.ID)
	if err != nil {
		return nil, err
	}

	value := new(domainpayment.Payment)
	value.ID = uuid.New()
	value.InvoiceID = request.InvoiceID
	value.Cryptogram = request.Cryptogram
	value.OrderID = request.OrderID
	value.CardSave = request.CardSave
	value.Name = request.Name
	err = u.paymentRepo.CreatePayment(ctx, value)
	if err != nil {
		return nil, err
	}
	order, err = u.orderRepo.GetOrderByID(ctx, request.OrderID)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func generateInvoiceID() string {
	return strconv.FormatInt(time.Now().UnixNano()%1_000_000_000_000_000, 10)
}

func GenerateCryptogram(publicKeyPEM []byte, payload CryptogramPayload) (string, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("marshal cryptogram payload: %w", err)
	}

	publicKey, err := parseRSAPublicKey(publicKeyPEM)
	if err != nil {
		return "", fmt.Errorf("parse rsa public key: %w", err)
	}

	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, body)
	if err != nil {
		return "", fmt.Errorf("encrypt cryptogram payload: %w", err)
	}

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func parseRSAPublicKey(publicKeyPEM []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(publicKeyPEM)
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	switch block.Type {
	case "PUBLIC KEY":
		key, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}

		rsaKey, ok := key.(*rsa.PublicKey)
		if !ok {
			return nil, errors.New("public key is not RSA")
		}

		return rsaKey, nil

	case "RSA PUBLIC KEY":
		return x509.ParsePKCS1PublicKey(block.Bytes)

	default:
		return nil, fmt.Errorf("unsupported public key type: %s", block.Type)
	}
}
