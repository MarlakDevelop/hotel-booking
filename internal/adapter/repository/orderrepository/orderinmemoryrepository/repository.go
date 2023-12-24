package orderinmemoryrepository

import "github.com/MarlakDevelop/hotel-booking/internal/domain/model"

type OrderInMemoryRepository struct {
	orders []*model.Order
}

func NewOrderInMemoryRepository() *OrderInMemoryRepository {
	return &OrderInMemoryRepository{orders: make([]*model.Order, 0)}
}
