package orderusecase

import (
	"github.com/MarlakDevelop/hotel-booking/internal/domain/monitor/logger"
	"github.com/MarlakDevelop/hotel-booking/internal/domain/repository/orderrepository"
	"github.com/MarlakDevelop/hotel-booking/internal/domain/repository/roomrepository"
)

type orderRepository interface {
	orderrepository.CreateOrder
	orderrepository.GetOrders
}

type roomRepository interface {
	roomrepository.GetRoom
}

type OrderUseCase struct {
	logger logger.WithContext

	orderRepository orderRepository
	roomRepository  roomRepository
}

func NewOrderUseCase(
	log logger.WithContext,
	orderRepository orderRepository,
	roomRepository roomRepository,
) *OrderUseCase {
	return &OrderUseCase{
		logger:          log,
		orderRepository: orderRepository,
		roomRepository:  roomRepository,
	}
}
