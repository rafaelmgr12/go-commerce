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
