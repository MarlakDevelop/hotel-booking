package orderrepository

import (
	"context"
	"time"
)

type CreateOrderIn struct {
	Room      string
	UserEmail string
	From      time.Time
	To        time.Time
}

type CreateOrderOut struct {
}

type CreateOrder interface {
	CreateOrder(ctx context.Context, in CreateOrderIn) (CreateOrderOut, error)
}
