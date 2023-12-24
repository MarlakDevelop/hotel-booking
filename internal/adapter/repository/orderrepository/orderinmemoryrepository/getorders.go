package orderinmemoryrepository

import (
	"context"

	pkgSlices "github.com/MarlakDevelop/hotel-booking/pkg/slices"

	"github.com/MarlakDevelop/hotel-booking/internal/domain/model"
	"github.com/MarlakDevelop/hotel-booking/internal/domain/repository/orderrepository"
)

func (repo *OrderInMemoryRepository) GetOrders(
	_ context.Context, in orderrepository.GetOrdersIn,
) (orderrepository.GetOrdersOut, error) {
	return orderrepository.GetOrdersOut{
		Orders: pkgSlices.Filter[*model.Order](repo.orders, repo.getOrdersFilterCB(in)),
	}, nil
}

func (repo *OrderInMemoryRepository) getOrdersFilterCB(
	in orderrepository.GetOrdersIn,
) func(order *model.Order, _ int) bool {
	const conditionsCount = 4

	conditionCBs := make([]func(*model.Order) bool, 0, conditionsCount)

	if in.Room != nil {
		conditionCBs = append(conditionCBs, func(order *model.Order) bool {
			return order.Room == *in.Room
		})
	}

	if in.UserEmail != nil {
		conditionCBs = append(conditionCBs, func(order *model.Order) bool {
			return order.UserEmail == *in.UserEmail
		})
	}

	if in.From != nil {
		conditionCBs = append(conditionCBs, func(order *model.Order) bool {
			return !order.To.Before(*in.From) //nolint:gocritic // in.From <= order.To.
		})
	}

	if in.To != nil {
		conditionCBs = append(conditionCBs, func(order *model.Order) bool {
			return !order.From.After(*in.To) //nolint:gocritic // in.To >= order.From.
		})
	}

	return func(order *model.Order, _ int) bool {
		for _, cb := range conditionCBs {
			if !cb(order) {
				return false
			}
		}

		return true
	}
}
