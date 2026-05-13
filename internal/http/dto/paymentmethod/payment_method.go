package paymentmethod

import "github.com/google/uuid"

type PaymentMethod struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Code string    `json:"code"`
}
