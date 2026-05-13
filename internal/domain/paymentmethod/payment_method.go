package paymentmethod

import "github.com/google/uuid"

type PaymentMethod struct {
	ID   uuid.UUID `gorm:"column:id"`
	Name string    `gorm:"column:name"`
	Code string    `gorm:"column:code"`
}

func (PaymentMethod) TableName() string {
	return "payment_method"
}
