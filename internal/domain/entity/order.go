package entity

import "github.com/google/uuid"

type Order struct {
	ID      uuid.UUID
	Product *Product
	Payment *Payment
	Labels  []string
}
