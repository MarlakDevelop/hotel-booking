package orderhttpcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MarlakDevelop/hotel-booking/internal/usecase/orderusecase"
)

func (cr *OrderHTTPController) GetOrders(res http.ResponseWriter, req *http.Request) {
	const logPrefix = "Controller GetOrders"

	ctx := req.Context()

	logger := cr.logger.WithContext(ctx)

	userEmail := req.URL.Query().Get("email")
	if userEmail == "" {
		logger.WarnKV(fmt.Sprintf("%s bad request", logPrefix), "userEmail", "field is nil")

		http.Error(res, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	out, err := cr.useCase.GetOrders(ctx, orderusecase.GetOrdersIn{UserEmail: userEmail})
	if err != nil {
		logger.ErrorF("%s useCase.GetOrders: %s", logPrefix, err.Error())

		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	content, err := json.Marshal(out.Orders)
	if err != nil {
		logger.ErrorF("%s json.Marshal: %s", logPrefix, err.Error())

		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	res.WriteHeader(http.StatusOK)

	_, err = res.Write(content)

	if err != nil {
		logger.ErrorF("%s res.Write: %s", logPrefix, err.Error())
	}

	logger.InfoF("%s finished successfully", logPrefix)
}
