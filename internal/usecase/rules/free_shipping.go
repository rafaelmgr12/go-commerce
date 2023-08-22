package rules

import "github.com/rafaelmgr12/go-commerce/internal/domain/entity"

type FreeShippingRule struct {
	next Rule
}

func (r *FreeShippingRule) SetNext(rule Rule) {
	r.next = rule
}

func (r *FreeShippingRule) Apply(o *entity.Order) {
	if o.Product.Value > 1000 {
		o.Labels = append(o.Labels, "frete-grátis")
	}
	if r.next != nil {
		r.next.Apply(o)
	}
}
