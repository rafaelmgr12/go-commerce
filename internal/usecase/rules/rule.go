package rules

import "github.com/rafaelmgr12/go-commerce/internal/domain/entity"

type Rule interface {
	SetNext(rule Rule)
	Apply(o *entity.Order)
}

type FreeShippingRule struct {
	next Rule
}

type FragileProductRule struct {
	next Rule
}

type ChildProductRule struct {
	next Rule
}

type BoletoDiscountRule struct {
	next Rule
}

func (r *FreeShippingRule) SetNext(rule Rule) {
	r.next = rule
}

func (r *FreeShippingRule) Apply(o *entity.Order) {
	if o.Product.Value > 1000 {
		o.Labels = append(o.Labels, "free-shipping")
	}
	if r.next != nil {
		r.next.Apply(o)
	}
}

func (r *FragileProductRule) SetNext(rule Rule) {
	r.next = rule
}

func (r *FragileProductRule) Apply(o *entity.Order) {
	if o.Product.Category == "home-appliance" {
		o.Labels = append(o.Labels, "fragile")
	}
	if r.next != nil {
		r.next.Apply(o)
	}
}

func (r *ChildProductRule) SetNext(rule Rule) {
	r.next = rule
}

func (r *ChildProductRule) Apply(o *entity.Order) {
	if o.Product.Category == "child" {
		o.Labels = append(o.Labels, "gift")
	}
	if r.next != nil {
		r.next.Apply(o)
	}
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
