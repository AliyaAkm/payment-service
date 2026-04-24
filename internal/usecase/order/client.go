package order

import (
	"context"
	"github.com/google/uuid"
	"payment-service/internal/domain/order"
	"payment-service/internal/domain/orderstatus"
	"payment-service/internal/domain/price"
)

type Repository interface {
	CreateOrder(ctx context.Context, value *order.Order) error
	GetOrderByID(ctx context.Context, id uuid.UUID) (*order.Order, error)
}
type StatusRepository interface {
	GetOrderStatusByCode(ctx context.Context, code string) (*orderstatus.OrderStatus, error)
}
type PriceRepository interface {
	GetCoursePriceByCourseID(ctx context.Context, courseID uuid.UUID) (*price.CoursePrice, error)
}
