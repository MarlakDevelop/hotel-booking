package orderrepository

import (
	"context"
	"time"

	"github.com/MarlakDevelop/hotel-booking/internal/domain/model"
)

type GetOrdersIn struct {
	Room      *string    // Filter by equal Room.
	UserEmail *string    // Filter by equal UserEmail.
	From      *time.Time // Filter by INTERSECTING From.
	To        *time.Time // Filter by INTERSECTING To.
}

type GetOrdersOut struct {
	Orders []*model.Order
}

type GetOrders interface {
	GetOrders(ctx context.Context, in GetOrdersIn) (GetOrdersOut, error)
}
