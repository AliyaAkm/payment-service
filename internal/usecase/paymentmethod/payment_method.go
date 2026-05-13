package paymentmethod

import (
	"context"
	"payment-service/internal/domain/paymentmethod"
)

func (u *UseCase) GetAllPaymentMethod(ctx context.Context) ([]paymentmethod.PaymentMethod, error) {
	return u.repo.GetAllPaymentMethod(ctx)
}
