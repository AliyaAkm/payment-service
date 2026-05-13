package payment

import (
	"context"
	"payment-service/internal/domain/payment"
)

func (r *Repo) CreatePayment(ctx context.Context, value *payment.Payment) error {
	err := r.db.WithContext(ctx).Create(value).Error
	if err != nil {
		return err
	}
	return nil
}
