package orderhttpcontroller

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	domainerror "github.com/MarlakDevelop/hotel-booking/internal/domain/error"
	"github.com/MarlakDevelop/hotel-booking/internal/usecase/orderusecase"
	"github.com/MarlakDevelop/hotel-booking/pkg/datetime"
)

type makeOrderRequestBody struct {
	Email string        `json:"email"`
	Room  string        `json:"room"`
	From  datetime.Date `json:"from"`
	To    datetime.Date `json:"to"`
}

func (cr *OrderHTTPController) MakeOrder(res http.ResponseWriter, req *http.Request) {
	const logPrefix = "Controller MakeOrder"

	ctx := req.Context()

	logger := cr.logger.WithContext(ctx)

	reqBody := new(makeOrderRequestBody)

	content, err := io.ReadAll(req.Body)

	if err != nil {
		logger.ErrorF("%s failed to read request body", logPrefix)

		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	err = json.Unmarshal(content, reqBody)

	if err != nil {
		logger.WarnF("%s can't unmarshal request body to json", logPrefix)

		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	_, err = cr.useCase.MakeOrder(ctx, orderusecase.MakeOrderIn{
		Room:      reqBody.Room,
		UserEmail: reqBody.Email,
		From:      reqBody.From.Time,
		To:        reqBody.To.Time,
	})

	if err != nil {
		cr.processMakeOrderError(ctx, res, err)
		return
	}

	res.WriteHeader(http.StatusCreated)

	logger.InfoF("%s finished successfully", logPrefix)
}

func (cr *OrderHTTPController) processMakeOrderError(ctx context.Context, res http.ResponseWriter, err error) {
	const logPrefix = "Controller MakeOrder"

	logger := cr.logger.WithContext(ctx)

	switch {
	case errors.Is(err, domainerror.ErrRoomNotFound), errors.Is(err, domainerror.ErrOrderTimeWindowConflict):
		logger.WarnF("%s bad request: %s", logPrefix, err.Error())

		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	case errors.Is(err, domainerror.ErrRoomAlreadyTaken):
		logger.WarnF("%s conflict: %s", logPrefix, err.Error())

		http.Error(res, http.StatusText(http.StatusConflict), http.StatusConflict)
	default:
		logger.ErrorF("%s internal error: %s", logPrefix, err.Error())

		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
