package paymentmethod

import (
	"context"
	"payment-service/internal/domain/paymentmethod"
)

type client interface {
	GetAllPaymentMethod(ctx context.Context) ([]paymentmethod.PaymentMethod, error)
}
