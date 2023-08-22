package rules

import "github.com/rafaelmgr12/go-commerce/internal/domain/entity"

type BoletoDiscountRule struct {
	next Rule
}

func (r *BoletoDiscountRule) SetNext(rule Rule) {
	r.next = rule
}

func (r *BoletoDiscountRule) Apply(o *entity.Order) {
	if o.Payment.PaymentMethod == "Boleto" {
		o.Payment.Value = int(float64(o.Payment.Value) * 0.9)
	}
	if r.next != nil {
		r.next.Apply(o)
	}
}
