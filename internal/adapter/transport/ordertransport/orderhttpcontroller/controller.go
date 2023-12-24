package orderhttpcontroller

import (
	"github.com/MarlakDevelop/hotel-booking/internal/domain/monitor/logger"
	"github.com/MarlakDevelop/hotel-booking/internal/usecase/orderusecase"
)

type OrderHTTPController struct {
	logger  logger.WithContext
	useCase *orderusecase.OrderUseCase
}

func NewOrderHTTPController(log logger.WithContext, useCase *orderusecase.OrderUseCase) *OrderHTTPController {
	return &OrderHTTPController{logger: log, useCase: useCase}
}
