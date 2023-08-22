package rules

import "github.com/rafaelmgr12/go-commerce/internal/domain/entity"

type ChildProductRule struct {
	next Rule
}

func (r *ChildProductRule) SetNext(rule Rule) {
	r.next = rule
}

func (r *ChildProductRule) Apply(o *entity.Order) {
	if o.Product.Category == "infantil" {
		o.Labels = append(o.Labels, "presente")
	}
	if r.next != nil {
		r.next.Apply(o)
	}
}
