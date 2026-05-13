package paymentmethod

import (
	"context"
	"payment-service/internal/domain/paymentmethod"
)

func (r *Repo) GetAllPaymentMethod(ctx context.Context) ([]paymentmethod.PaymentMethod, error) {
	var paymentMethod []paymentmethod.PaymentMethod
	err := r.db.WithContext(ctx).Find(&paymentMethod).Error
	if err != nil {
		return nil, err
	}
	return paymentMethod, nil
}
