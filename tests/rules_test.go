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

func TestNewProduct(t *testing.T) {
	product, err := entity.NewProduct("electronics", 500)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, "electronics", product.Category, "Product category mismatch")
	assert.Equal(t, 500, product.Value, "Product value mismatch")
}

func TestNewPader(t *testing.T) {
	payment, err := entity.NewPader("CreditCard", 500)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, "CreditCard", payment.PaymentMethod, "Payment method mismatch")
	assert.Equal(t, 500, payment.Value, "Payment value mismatch")
}

func TestSingleRuleApplication(t *testing.T) {
	// Create an order that should only trigger the FreeShippingRule
	order := &entity.Order{
		Product: &entity.Product{
			Value: 1100, // Value > 1000 should trigger FreeShippingRule
		},
		Payment: &entity.Payment{
			PaymentMethod: "CreditCard", // This should not trigger BoletoDiscountRule
			Value:         1100,
		},
	}

	// Declare all rules
	ruleManager := rules.NewRuleManager()
	ruleManager.AddRule(&rules.FreeShippingRule{})
	ruleManager.AddRule(&rules.FragileProductRule{})
	ruleManager.AddRule(&rules.ChildProductRule{})
	ruleManager.AddRule(&rules.BoletoDiscountRule{})

	// Apply rules
	ruleManager.ApplyRules(order)

	// Assertions
	assert.Contains(t, order.Labels, "frete-grátis", "FreeShippingRule was not applied correctly")
	assert.NotContains(t, order.Labels, "frágil", "FragileProductRule was incorrectly applied")
	assert.NotContains(t, order.Labels, "presente", "ChildProductRule was incorrectly applied")
	assert.Equal(t, 1100, order.Payment.Value, "BoletoDiscountRule was incorrectly applied")
}

func TestRuleManager(t *testing.T) {
	ruleManager := rules.NewRuleManager()
	assert.Equal(t, 0, len(ruleManager.GetRules()), "Initial rules should be empty")

	// Add rules
	ruleManager.AddRule(&rules.FreeShippingRule{})
	ruleManager.AddRule(&rules.FragileProductRule{})
	assert.Equal(t, 2, len(ruleManager.GetRules()), "Rules count mismatch after adding")

	// Remove rule
	ruleManager.RemoveRule(&rules.FreeShippingRule{})
	assert.Equal(t, 1, len(ruleManager.GetRules()), "Rules count mismatch after removing")

	// Apply rules
	order := &entity.Order{
		Product: &entity.Product{
			Category: "eletrodómestico",
			Value:    500,
		},
	}
	ruleManager.ApplyRules(order)
	assert.Contains(t, order.Labels, "frágil", "FragileProductRule was not applied correctly")
}
