package entity

import "github.com/google/uuid"

type Product struct {
	ID       uuid.UUID
	Category string
	Value    int
}

func NewProduct(category string, value int) (*Product, error) {
	return &Product{
		ID:       uuid.New(),
		Category: category,
		Value:    value}, nil
}
