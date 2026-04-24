package paymentstatus

import (
	"context"
	"github.com/google/uuid"
	"payment-service/internal/domain/paymentstatus"
)

func (r *Repo) GetPaymentStatusByID(ctx context.Context, id uuid.UUID) (*paymentstatus.PaymentStatus, error) {
	var entity *paymentstatus.PaymentStatus
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *Repo) GetPaymentStatusByCode(ctx context.Context, code string) (*paymentstatus.PaymentStatus, error) {
	var entity *paymentstatus.PaymentStatus
	err := r.db.WithContext(ctx).Where("code = ?", code).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}
