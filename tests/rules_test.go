package tests

import (
	"testing"

	"github.com/rafaelmgr12/go-commerce/internal/domain/entity"
	"github.com/rafaelmgr12/go-commerce/internal/usecase/rules"
	"github.com/stretchr/testify/assert"
)

func TestFreeShippingRule(t *testing.T) {
	order := &entity.Order{
		Product: &entity.Product{
			Value: 1100,
		},
	}

	rule := &rules.FreeShippingRule{}
	rule.Apply(order)

	assert.Contains(t, order.Labels, "free-shipping", "FreeShippingRule was not applied correctly")
}

func TestFragileProductRule(t *testing.T) {
	order := &entity.Order{
		Product: &entity.Product{
			Category: "home-appliance",
		},
	}

	rule := &rules.FragileProductRule{}
	rule.Apply(order)

	assert.Contains(t, order.Labels, "fragile", "FragileProductRule was not applied correctly")
}

func TestChildProductRule(t *testing.T) {
	order := &entity.Order{
		Product: &entity.Product{
			Category: "child",
		},
	}

	rule := &rules.ChildProductRule{}
	rule.Apply(order)

	assert.Contains(t, order.Labels, "gift", "ChildProductRule was not applied correctly")
}

func TestBoletoDiscountRule(t *testing.T) {
	order := &entity.Order{
		Product: &entity.Product{
			Value: 1000,
		},
		Payment: &entity.Payment{
			PaymentMethod: "Boleto",
			Value:         1000,
		},
	}

	rule := &rules.BoletoDiscountRule{}
	rule.Apply(order)

	assert.Equal(t, 900, order.Payment.Value, "BoletoDiscountRule was not applied correctly")
}

func TestChainOfRules(t *testing.T) {
	order := &entity.Order{
		Product: &entity.Product{
			Category: "home-appliance",
			Value:    1100,
		},
		Payment: &entity.Payment{
			PaymentMethod: "Boleto",
			Value:         1100,
		},
	}

	r1 := &rules.FreeShippingRule{}
	r2 := &rules.FragileProductRule{}
	r3 := &rules.ChildProductRule{}
	r4 := &rules.BoletoDiscountRule{}

	r1.SetNext(r2)
	r2.SetNext(r3)
	r3.SetNext(r4)

	r1.Apply(order)

	assert.Equal(t, 2, len(order.Labels), "Chain of rules failed")
	assert.Equal(t, 990, order.Payment.Value, "Chain of rules failed")
}
