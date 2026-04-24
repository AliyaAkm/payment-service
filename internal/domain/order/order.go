package order

import (
	"github.com/google/uuid"
	"payment-service/internal/domain/orderstatus"
)

type Order struct {
	ID       uuid.UUID `gorm:"column:id"`
	UserID   uuid.UUID `gorm:"column:user_id"`
	CourseID uuid.UUID `gorm:"column:course_id"`
	Amount   int       `gorm:"column:amount"`
	Currency string    `gorm:"column:currency"`
	StatusID uuid.UUID `gorm:"column:status_id"`

	Status orderstatus.OrderStatus
}

func (Order) TableName() string {
	return "orders"
}
