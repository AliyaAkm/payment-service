package orderstatus

import (
	"context"
	"github.com/google/uuid"
	"payment-service/internal/domain/orderstatus"
)

func (r *Repo) GetOrderStatusByID(ctx context.Context, id uuid.UUID) (*orderstatus.OrderStatus, error) {
	var entity *orderstatus.OrderStatus
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *Repo) GetOrderStatusByCode(ctx context.Context, code string) (*orderstatus.OrderStatus, error) {
	var entity *orderstatus.OrderStatus
	err := r.db.WithContext(ctx).Where("code = ?", code).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}
