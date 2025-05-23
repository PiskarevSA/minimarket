package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	json "github.com/bytedance/sonic"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/entities"
)

type GetOrdersUsecase interface {
	Do(
		ctx context.Context,
		userId uuid.UUID,
		offset,
		limit int32,
	) (orders []entities.Order, err error)
}

type GetOrders struct {
	usecase GetOrdersUsecase
}

func NewGetOrders(usecase GetOrdersUsecase) *GetOrders {
	return &GetOrders{usecase: usecase}
}

func (h *GetOrders) Mount(r chi.Router) {
	r.Get("/api/user/orders", h.handle)
}

type OrderItem struct {
	Number     string    `json:"number"`
	Status     string    `json:"status"`
	Accrual    string    `json:"accrual"`
	UploadedAt time.Time `json:"uploadedAt"`
}

type GetOrdersResponse []OrderItem

func (h *GetOrders) handle(rw http.ResponseWriter, req *http.Request) {
	const op = "getOrders"

	ctx := req.Context()

	token, ok := getJwtFromContext(ctx, op)
	if !ok {
		writeInternalServerError(rw)

		return
	}

	userId, ok := getUserIdFromJwt(token, op)
	if !ok {
		writeInternalServerError(rw)

		return
	}

	offset, limit, ok := h.getPaginationFromQuery(rw, req)
	if !ok {
		return
	}

	orders, err := h.usecase.Do(ctx, userId, offset, limit)
	if err != nil {
		writeInternalServerError(rw)

		return
	}

	if len(orders) == 0 {
		rw.WriteHeader(http.StatusNoContent)

		return
	}

	respData := h.ordersToGetOrdersResponse(orders)

	rw.WriteHeader(http.StatusOK)
	enc := json.ConfigDefault.NewEncoder(rw)
	enc.Encode(respData)
}

func (h *GetOrders) getPaginationFromQuery(
	rw http.ResponseWriter,
	req *http.Request,
) (offset, limit int32, ok bool) {
	const (
		minOffset int32 = 0
		minLimit  int32 = 1
		maxLimit  int32 = 35
	)

	offset = int32(0)
	limit = int32(10)

	offsetParam := req.URL.Query().Get("offset")
	if offsetParam != "" {
		parsed, parseErr := strconv.ParseInt(offsetParam, 10, 32)
		if parseErr != nil {
			writeValidationError(
				rw,
				http.StatusBadRequest,
				"V1012",
				"offset",
				"invalid integer",
			)

			return offset, limit, false
		}

		offset = int32(parsed)
		if offset < minOffset {
			writeValidationError(
				rw,
				http.StatusBadRequest,
				"V1012",
				"offset",
				"invalid integer",
			)

			return offset, limit, false
		}
	}

	limitParam := req.URL.Query().Get("limit")
	if limitParam != "" {
		parsed, parseErr := strconv.ParseInt(limitParam, 10, 32)
		if parseErr != nil {
			writeValidationError(
				rw,
				http.StatusBadRequest,
				"V1012",
				"limit",
				"invalid integer",
			)

			return offset, limit, false
		}

		limit = int32(parsed)
		if limit < minLimit || limit > maxLimit {
			writeValidationError(
				rw,
				http.StatusBadRequest,
				"V1012",
				"offset",
				"invalid integer",
			)

			return offset, limit, false
		}
	}

	return offset, limit, true
}

func (h *GetOrders) ordersToGetOrdersResponse(
	orders []entities.Order,
) GetOrdersResponse {
	respData := make(GetOrdersResponse, len(orders))

	for i, order := range orders {
		respData[i].Number = order.Number().String()
		respData[i].Status = order.Status().String()
		respData[i].Accrual = order.Accrual().String()
		respData[i].UploadedAt = order.UploadedAt()
	}

	return respData
}
