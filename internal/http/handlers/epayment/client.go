package epayment

import (
	"context"
	"payment-service/internal/domain/order"
	"payment-service/service/epayment"
)

type client interface {
	GetToken(ctx context.Context, request epayment.GetTokenRequest) (epayment.GetTokenResponse, error)
	GetStatusTransaction(ctx context.Context, invoiceID string) (epayment.GetTransactionResponse, error)
}
type paymentUseCase interface {
	CreatePayment(ctx context.Context, request epayment.PaymentRequest) (*order.Order, error)
}
