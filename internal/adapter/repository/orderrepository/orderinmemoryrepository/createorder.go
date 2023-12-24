package orderinmemoryrepository

import (
	"context"

	"github.com/MarlakDevelop/hotel-booking/internal/domain/model"
	"github.com/MarlakDevelop/hotel-booking/internal/domain/repository/orderrepository"
)

func (repo *OrderInMemoryRepository) CreateOrder(
	_ context.Context, in orderrepository.CreateOrderIn,
) (orderrepository.CreateOrderOut, error) {
	order := &model.Order{
		Room:      in.Room,
		UserEmail: in.UserEmail,
		From:      in.From,
		To:        in.To,
	}

	repo.orders = append(repo.orders, order)

	return orderrepository.CreateOrderOut{}, nil
}
