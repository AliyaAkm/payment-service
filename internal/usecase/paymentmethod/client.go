package paymentmethod

import (
	"context"
	"payment-service/internal/domain/paymentmethod"
)

type Repository interface {
	GetAllPaymentMethod(ctx context.Context) ([]paymentmethod.PaymentMethod, error)
}
