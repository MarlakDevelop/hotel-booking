package orderusecase

import (
	"context"
	"fmt"

	"github.com/MarlakDevelop/hotel-booking/internal/domain/model"
	"github.com/MarlakDevelop/hotel-booking/internal/domain/repository/orderrepository"
)

type GetOrdersIn struct {
	UserEmail string
}

type GetOrdersOut struct {
	Orders []*model.Order
}

func (uc *OrderUseCase) GetOrders(ctx context.Context, in GetOrdersIn) (GetOrdersOut, error) {
	logger := uc.logger.WithContext(ctx)

	logger.DebugKV("UseCase GetOrders started", "userEmail", in.UserEmail)

	getOrdersOut, err := uc.orderRepository.GetOrders(ctx, orderrepository.GetOrdersIn{UserEmail: &in.UserEmail})

	if err != nil {
		logger.DebugF("Orders getting: %s", err.Error())

		return GetOrdersOut{}, fmt.Errorf("orderRepository.GetOrders: %w", err)
	}

	logger.DebugKV("UseCase MakeOrder successfully finished", "orders", getOrdersOut.Orders)

	return GetOrdersOut{Orders: getOrdersOut.Orders}, nil
}
