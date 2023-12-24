package app

import (
	"fmt"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/MarlakDevelop/hotel-booking/internal/adapter/transport/ordertransport/orderhttpcontroller"
	"github.com/MarlakDevelop/hotel-booking/internal/infrastructure/config"
)

func runHTTPServer(cfg *config.Config, orderHTTPController *orderhttpcontroller.OrderHTTPController) error {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(cfg.HTTP.Timeout))

	router.Get("/orders", orderHTTPController.GetOrders)
	router.Post("/order", orderHTTPController.MakeOrder)

	//nolint:gosec // using timeout middleware instead.
	err := http.ListenAndServe(net.JoinHostPort("", cfg.HTTP.Port), router)
	if err != nil {
		return fmt.Errorf("http.ListenAndServe: %w", err)
	}

	return nil
}
