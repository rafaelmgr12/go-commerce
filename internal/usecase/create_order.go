package usecase

import (
	"github.com/rafaelmgr12/go-commerce/internal/domain/entity"
	"github.com/rafaelmgr12/go-commerce/internal/usecase/rules"
)

var ruleManager *rules.RuleManager

func init() {
	ruleManager = rules.NewRuleManager()
	ruleManager.AddRule(&rules.FreeShippingRule{})
	ruleManager.AddRule(&rules.FragileProductRule{})
	ruleManager.AddRule(&rules.ChildProductRule{})
	ruleManager.AddRule(&rules.BoletoDiscountRule{})
}

func ProcessOrder(order *entity.Order) {
	ruleManager.ApplyRules(order)
}
