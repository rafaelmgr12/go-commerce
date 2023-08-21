package entity

import "github.com/google/uuid"

type Product struct {
	ID       uuid.UUID
	Category string
	Value    int
}
