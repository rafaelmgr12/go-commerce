package usecase

import (
	"github.com/rafaelmgr12/go-commerce/internal/domain/entity"
	"github.com/rafaelmgr12/go-commerce/internal/usecase/rules"
)

func ProcessOrder(order *entity.Order, startingRule rules.Rule) {
	startingRule.Apply(order)
}
