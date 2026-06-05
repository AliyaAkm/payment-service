package payment

import (
	"context"
	"github.com/google/uuid"
	"payment-service/internal/domain/order"
	"payment-service/internal/domain/orderstatus"
	"payment-service/internal/domain/payment"
	"payment-service/internal/domain/subscription"
	paymentclient "payment-service/service/epayment"
)

type OrderRepository interface {
	GetOrderByID(ctx context.Context, id uuid.UUID) (*order.Order, error)
	UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, statusID uuid.UUID) error
}
type PaymentRepo interface {
	CreatePayment(ctx context.Context, value *payment.Payment) error
}
type PaymentClient interface {
	CreatePayment(ctx context.Context, request paymentclient.PaymentRequest) (paymentclient.PaymentResponse, error)
	GetStatusTransaction(ctx context.Context, invoiceID string) (paymentclient.GetTransactionResponse, error)
}
type OrderStatusRepo interface {
	GetOrderStatusByCode(ctx context.Context, code string) (*orderstatus.OrderStatus, error)
}
type SubscriptionRepo interface {
	CreateSubscription(ctx context.Context, value *subscription.Subscription) error
	GetSubscriptionByID(ctx context.Context, id uuid.UUID) (*subscription.Subscription, error)
}

type NotificationSender interface {
	SendEvent(ctx context.Context, userID uuid.UUID, event string, data map[string]any) error
}
