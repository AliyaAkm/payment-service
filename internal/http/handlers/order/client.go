package order

import (
	"context"
	"payment-service/internal/domain/order"
)

type client interface {
	CreateOrder(ctx context.Context, value *order.Order) (*order.Order, error)
}
