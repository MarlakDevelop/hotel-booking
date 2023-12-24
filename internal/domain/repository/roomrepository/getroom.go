package roomrepository

import (
	"context"

	"github.com/MarlakDevelop/hotel-booking/internal/domain/model"
)

type GetRoomIn struct {
	Name string
}

type GetRoomOut struct {
	Room *model.Room
}

type GetRoom interface {
	GetRoom(ctx context.Context, in GetRoomIn) (GetRoomOut, error)
}
