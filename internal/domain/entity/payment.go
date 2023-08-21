package entity

import "github.com/google/uuid"

type Payment struct {
	ID            uuid.UUID
	PaymentMethod string
	Value         int
}

func NewPader(payment string, value int) (*Payment, error) {
	return &Payment{
		ID:            uuid.New(),
		PaymentMethod: payment,
		Value:         value}, nil
}
