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

	ruleManager := rules.NewRuleManager()
	ruleManager.AddRule(&rules.FreeShippingRule{})
	ruleManager.ApplyRules(order)

	assert.Contains(t, order.Labels, "frete-grátis", "FreeShippingRule was not applied correctly")
}

func TestFragileProductRule(t *testing.T) {
	order := &entity.Order{
		Product: &entity.Product{
			Category: "eletrodómestico",
		},
	}

	ruleManager := rules.NewRuleManager()
	ruleManager.AddRule(&rules.FragileProductRule{})
	ruleManager.ApplyRules(order)

	assert.Contains(t, order.Labels, "frágil", "FragileProductRule was not applied correctly")
}

func TestChildProductRule(t *testing.T) {
	order := &entity.Order{
		Product: &entity.Product{
			Category: "infantil",
		},
	}

	ruleManager := rules.NewRuleManager()
	ruleManager.AddRule(&rules.ChildProductRule{})
	ruleManager.ApplyRules(order)

	assert.Contains(t, order.Labels, "presente", "ChildProductRule was not applied correctly")
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

	ruleManager := rules.NewRuleManager()
	ruleManager.AddRule(&rules.BoletoDiscountRule{})
	ruleManager.ApplyRules(order)

	assert.Equal(t, 900, order.Payment.Value, "BoletoDiscountRule was not applied correctly")
}

func TestChainOfRules(t *testing.T) {
	order := &entity.Order{
		Product: &entity.Product{
			Category: "eletrodómestico",
			Value:    1100,
		},
		Payment: &entity.Payment{
			PaymentMethod: "Boleto",
			Value:         1100,
		},
	}

	ruleManager := rules.NewRuleManager()
	ruleManager.AddRule(&rules.FreeShippingRule{})
	ruleManager.AddRule(&rules.FragileProductRule{})
	ruleManager.AddRule(&rules.ChildProductRule{})
	ruleManager.AddRule(&rules.BoletoDiscountRule{})
	ruleManager.ApplyRules(order)

	assert.Equal(t, 2, len(order.Labels), "Chain of rules failed")
	assert.Equal(t, 990, order.Payment.Value, "Chain of rules failed")
}
