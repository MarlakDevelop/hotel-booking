package domainerror

import "errors"

var (
	ErrOrderTimeWindowConflict = errors.New("order time window conflict")
)
