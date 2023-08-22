package rules

import "github.com/rafaelmgr12/go-commerce/internal/domain/entity"

type Rule interface {
	SetNext(rule Rule)
	Apply(o *entity.Order)
}

type RuleManager struct {
	rules []Rule
}

func NewRuleManager() *RuleManager {
	return &RuleManager{}
}

func (rm *RuleManager) AddRule(rule Rule) {
	rm.rules = append(rm.rules, rule)
}

func (rm *RuleManager) RemoveRule(rule Rule) {
	// Logic to remove a rule
	for i, r := range rm.rules {
		if r == rule {
			rm.rules = append(rm.rules[:i], rm.rules[i+1:]...)
			break
		}
	}
}

func (rm *RuleManager) GetRules() []Rule {
	return rm.rules
}

func (rm *RuleManager) ApplyRules(o *entity.Order) {
	for _, rule := range rm.rules {
		rule.Apply(o)
	}
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
		o.Labels = append(o.Labels, "frete-grátis")
	}
	if r.next != nil {
		r.next.Apply(o)
	}
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
