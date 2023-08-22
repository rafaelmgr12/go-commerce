package rules

import (
	"reflect"

	"github.com/rafaelmgr12/go-commerce/internal/domain/entity"
)

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

func (rm *RuleManager) RemoveRule(ruleToRemove Rule) {
	var newRules []Rule
	for _, rule := range rm.rules {
		if reflect.TypeOf(rule) != reflect.TypeOf(ruleToRemove) {
			newRules = append(newRules, rule)
		}
	}
	rm.rules = newRules
}

func (rm *RuleManager) GetRules() []Rule {
	return rm.rules
}

func (rm *RuleManager) ApplyRules(o *entity.Order) {
	for _, rule := range rm.rules {
		rule.Apply(o)
	}
}
