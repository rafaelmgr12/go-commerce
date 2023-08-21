package entity

import "github.com/google/uuid"

type Order struct {
	ID      uuid.UUID
	Product *Product
	Payment *Payment
	Labels  []string
}

func NewOrder(product *Product, payment *Payment, labels []string) (*Order, error) {
	return &Order{
		ID:      uuid.New(),
		Product: product,
		Payment: payment,
		Labels:  labels,
	}, nil
}
