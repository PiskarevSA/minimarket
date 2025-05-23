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

type GetWithdrawalsUsecase interface {
	Do(
		ctx context.Context,
		userId uuid.UUID,
		offset, limit int32,
	) (txs []entities.Transaction, err error)
}

type GetWithdrawals struct {
	usecase GetWithdrawalsUsecase
}

func NewGetWithdrawals(usecase GetWithdrawalsUsecase) *GetWithdrawals {
	return &GetWithdrawals{usecase: usecase}
}

func (h *GetWithdrawals) Mount(r chi.Router) {
	r.Get("/api/user/withdrawals", h.handle)
}

type WithdrawalItem struct {
	Order       string    `json:"order"`
	Sum         string    `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}

type GetWithdrawalsResponse []WithdrawalItem

func (h *GetWithdrawals) handle(rw http.ResponseWriter, req *http.Request) {
	const op = "getWithdrawals"

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

	withdrawals, err := h.usecase.Do(ctx, userId, offset, limit)
	if err != nil {
		writeInternalServerError(rw)

		return
	}

	if len(withdrawals) == 0 {
		rw.WriteHeader(http.StatusNoContent)

		return
	}

	respData := h.withdrawalsToGetWithdrawalsResponse(withdrawals)

	rw.WriteHeader(http.StatusOK)
	enc := json.ConfigDefault.NewEncoder(rw)
	enc.Encode(respData)
}

func (h *GetWithdrawals) getPaginationFromQuery(
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

func (h *GetWithdrawals) withdrawalsToGetWithdrawalsResponse(
	withdrawals []entities.Transaction,
) GetWithdrawalsResponse {
	resp := make(GetWithdrawalsResponse, len(withdrawals))

	for i, tx := range withdrawals {
		resp[i].Order = tx.OrderNumber().String()
		resp[i].Sum = tx.Sum().String()
		resp[i].ProcessedAt = tx.ProcessedAt()
	}

	return resp
}
