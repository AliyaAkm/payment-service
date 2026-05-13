package payment

import "github.com/google/uuid"

type Payment struct {
	ID         uuid.UUID `gorm:"column:id"`
	InvoiceID  string    `gorm:"column:invoice_id"`
	Cryptogram string    `gorm:"column:cryptogram"`
	OrderID    uuid.UUID `gorm:"column:order_id"`
	CardSave   bool      `gorm:"column:card_save"`
	Name       string    `gorm:"column:name"`
}

func (Payment) TableName() string {
	return "payments"
}
