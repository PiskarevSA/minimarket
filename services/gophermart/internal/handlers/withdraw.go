package handlers

import (
	"context"
	"errors"
	"net/http"

	json "github.com/bytedance/sonic"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/usecases"
)

type WithdrawUsecase interface {
	Do(
		ctx context.Context,
		userId uuid.UUID,
		orderNumber,
		sum string,
	) error
}

type Withdraw struct {
	usecase WithdrawUsecase
}

func NewWithdraw(usecase WithdrawUsecase) *Withdraw {
	return &Withdraw{usecase: usecase}
}

func (h *Withdraw) Mount(r chi.Router) {
	r.Post("/api/user/balance/withdraw", h.handle)
}

type WithdrawRequest struct {
	Order string `json:"order"`
	Sum   string `json:"sum"`
}

func (h *Withdraw) handle(rw http.ResponseWriter, req *http.Request) {
	const op = "withdraw"

	ctx := req.Context()

	token, ok := getJwtFromContext(ctx, "withdraw")
	if !ok {
		writeInternalServerError(rw)

		return
	}

	userId, ok := getUserIdFromJwt(token, "withdraw")
	if !ok {
		writeInternalServerError(rw)

		return
	}

	var reqData WithdrawRequest

	dec := json.ConfigDefault.NewDecoder(req.Body)

	err := dec.Decode(&reqData)
	if err != nil {
		writeValidationError(
			rw,
			http.StatusBadRequest,
			"V1042",
			"body",
			"invalid json format",
		)

		return
	}

	err = h.usecase.Do(ctx, userId, reqData.Order, reqData.Sum)
	if err != nil {
		{
			var e *usecases.ValidationError

			var e1 *usecases.BusinessError

			switch {
			case errors.As(err, &e):
				statusCode := http.StatusBadRequest

				if e.IsCodeMatch("V1107") {
					statusCode = http.StatusUnprocessableEntity
				}

				writeValidationError(
					rw,
					statusCode,
					e.Code,
					e.Field,
					e.Message,
				)
			case errors.As(err, &e1):
				writeBusinessError(
					rw,
					http.StatusPaymentRequired,
					e1.Code,
					e1.Message,
				)
			default:
				writeInternalServerError(rw)
			}

			return
		}
	}

	rw.WriteHeader(http.StatusOK)
}
