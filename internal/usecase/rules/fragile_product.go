package rules

import "github.com/rafaelmgr12/go-commerce/internal/domain/entity"

type FragileProductRule struct {
	next Rule
}

func (r *FragileProductRule) SetNext(rule Rule) {
	r.next = rule
}

func (r *FragileProductRule) Apply(o *entity.Order) {
	if o.Product.Category == "eletrodómestico" {
		o.Labels = append(o.Labels, "frágil")
	}
	if r.next != nil {
		r.next.Apply(o)
	}
}
