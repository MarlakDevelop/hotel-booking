package roominmemoryrepository

import "github.com/MarlakDevelop/hotel-booking/internal/domain/model"

type RoomInMemoryRepository struct {
	rooms map[string]*model.Room
}

func NewRoomInMemoryRepository(
	rooms []*model.Room,
) *RoomInMemoryRepository {
	hashedRooms := make(map[string]*model.Room, len(rooms))
	for _, room := range rooms {
		hashedRooms[room.Name] = room
	}

	return &RoomInMemoryRepository{rooms: hashedRooms}
}
