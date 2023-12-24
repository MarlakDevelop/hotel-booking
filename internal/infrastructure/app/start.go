package app

import (
	"context"
	"fmt"
	"os"

	"github.com/MarlakDevelop/hotel-booking/internal/adapter/repository/orderrepository/orderinmemoryrepository"
	"github.com/MarlakDevelop/hotel-booking/internal/adapter/repository/roomrepository/roominmemoryrepository"
	"github.com/MarlakDevelop/hotel-booking/internal/adapter/transport/ordertransport/orderhttpcontroller"
	"github.com/MarlakDevelop/hotel-booking/internal/domain/model"
	"github.com/MarlakDevelop/hotel-booking/internal/infrastructure/config"
	"github.com/MarlakDevelop/hotel-booking/internal/infrastructure/monitor"
	"github.com/MarlakDevelop/hotel-booking/internal/usecase/orderusecase"
)

func Start(_ context.Context) error {
	cfg := config.NewConfig()

	logger := monitor.NewLogger(os.Stdout, cfg.Logger.Level)

	logger.DebugF("cfg: %v", cfg)

	roomRepository := roominmemoryrepository.NewRoomInMemoryRepository(
		[]*model.Room{{Name: "econom"}, {Name: "standart"}, {Name: "lux"}},
	)
	orderRepository := orderinmemoryrepository.NewOrderInMemoryRepository()

	orderUseCase := orderusecase.NewOrderUseCase(logger, orderRepository, roomRepository)

	orderController := orderhttpcontroller.NewOrderHTTPController(logger, orderUseCase)

	logger.InfoF("Service start")

	err := runHTTPServer(cfg, orderController)
	if err != nil {
		return fmt.Errorf("runHTTPServer: %w", err)
	}

	return nil
}
