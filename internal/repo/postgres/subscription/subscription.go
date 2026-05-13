package subscription

import (
	"context"
	"github.com/google/uuid"
	"payment-service/internal/domain/subscription"
)

func (r *Repo) CreateSubscription(ctx context.Context, value *subscription.Subscription) error {
	err := r.db.WithContext(ctx).Create(value).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetSubscriptionByID(ctx context.Context, id uuid.UUID) (*subscription.Subscription, error) {
	var entity subscription.Subscription
	err := r.db.WithContext(ctx).First(&entity, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
