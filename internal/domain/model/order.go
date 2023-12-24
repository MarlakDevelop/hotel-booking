package model

import "time"

type Order struct {
	Room      string
	UserEmail string
	From      time.Time
	To        time.Time
}
