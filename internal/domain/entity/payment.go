package entity

import "github.com/google/uuid"

type Payment struct {
	ID            uuid.UUID
	PaymentMethod string
	Value         int
}
