package orderusecase

import (
	"context"
	"fmt"
	"time"

	domainError "github.com/MarlakDevelop/hotel-booking/internal/domain/error"
	"github.com/MarlakDevelop/hotel-booking/internal/domain/repository/orderrepository"
	"github.com/MarlakDevelop/hotel-booking/internal/domain/repository/roomrepository"
)

type MakeOrderIn struct {
	Room      string
	UserEmail string
	From      time.Time
	To        time.Time
}

type MakeOrderOut struct {
}

func (uc *OrderUseCase) MakeOrder(ctx context.Context, in MakeOrderIn) (MakeOrderOut, error) {
	logger := uc.logger.WithContext(ctx)

	logger.DebugKV(
		"UseCase MakeOrder started",
		"room", in.Room,
		"userEmail", in.UserEmail,
		"from", in.From,
		"to", in.To,
	)

	if in.From.After(in.To) {
		logger.DebugKV("Order time window conflict found", "from", in.From, "to", in.To)

		return MakeOrderOut{}, domainError.ErrOrderTimeWindowConflict
	}

	_, err := uc.roomRepository.GetRoom(ctx, roomrepository.GetRoomIn{Name: in.Room})

	if err != nil {
		logger.DebugF("Room getting: %s", err.Error())

		return MakeOrderOut{}, fmt.Errorf("roomRepository.GetRoom: %w", err)
	}

	var getOrdersOut orderrepository.GetOrdersOut

	getOrdersOut, err = uc.orderRepository.GetOrders(ctx, orderrepository.GetOrdersIn{
		Room: &in.Room,
		From: &in.From,
		To:   &in.To,
	})

	if err != nil {
		logger.DebugF("Orders getting: %s", err.Error())

		return MakeOrderOut{}, fmt.Errorf("orderRepository.GetOrders: %w", err)
	}

	if len(getOrdersOut.Orders) > 0 {
		logger.DebugKV("Orders found, but mustn't", "orders", getOrdersOut.Orders)

		return MakeOrderOut{}, domainError.ErrRoomAlreadyTaken
	}

	_, err = uc.orderRepository.CreateOrder(ctx, orderrepository.CreateOrderIn{
		Room:      in.Room,
		UserEmail: in.UserEmail,
		From:      in.From,
		To:        in.To,
	})

	if err != nil {
		logger.DebugF("Order creation: %s", err.Error())

		return MakeOrderOut{}, fmt.Errorf("orderRepository.CreateOrder: %w", err)
	}

	logger.DebugF("UseCase MakeOrder successfully finished")

	return MakeOrderOut{}, nil
}
