package order

import (
	"github.com/google/uuid"
	"payment-service/internal/domain/orderstatus"
)

type OrderRequest struct {
	UserID   uuid.UUID `json:"user_id"`
	CourseID uuid.UUID `json:"course_id"`
	Amount   int       `json:"amount"`
	Currency string    `json:"currency"`
}

type OrderResponse struct {
	ID       uuid.UUID `json:"id"`
	UserID   uuid.UUID `json:"user_id"`
	CourseID uuid.UUID `json:"course_id"`
	Amount   int       `json:"amount"`
	Currency string    `json:"currency"`

	Status orderstatus.OrderStatus `json:"status"`
}
