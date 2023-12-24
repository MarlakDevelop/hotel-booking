package roominmemoryrepository

import (
	"context"

	domainerror "github.com/MarlakDevelop/hotel-booking/internal/domain/error"
	"github.com/MarlakDevelop/hotel-booking/internal/domain/repository/roomrepository"
)

func (repo *RoomInMemoryRepository) GetRoom(
	_ context.Context, in roomrepository.GetRoomIn,
) (roomrepository.GetRoomOut, error) {
	room, roomFound := repo.rooms[in.Name]
	if !roomFound {
		return roomrepository.GetRoomOut{}, domainerror.ErrRoomNotFound
	}

	return roomrepository.GetRoomOut{Room: room}, nil
}
