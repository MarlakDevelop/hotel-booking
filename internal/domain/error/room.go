package domainerror

import "errors"

var (
	ErrRoomNotFound     = errors.New("room not found")
	ErrRoomAlreadyTaken = errors.New("room already taken on this time")
)
