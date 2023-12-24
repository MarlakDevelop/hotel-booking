package main

import (
	"context"
	"log"

	"github.com/MarlakDevelop/hotel-booking/internal/infrastructure/app"
)

func main() {
	if err := app.Start(context.Background()); err != nil {
		log.Fatalf("Service fatal error: %s", err.Error())
	}
}
