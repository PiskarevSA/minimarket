package handlers

import (
	"context"
	"errors"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/usecases"
)

type UploadOrderNumberUsecase interface {
	Do(
		ctx context.Context,
		userID uuid.UUID,
		orderNumber string,
	) error
}

type UploadOrderNumber struct {
	usecase UploadOrderNumberUsecase
}

func NewUploadOrderNumber(usecase UploadOrderNumberUsecase) *UploadOrderNumber {
	return &UploadOrderNumber{usecase: usecase}
}

func (h *UploadOrderNumber) Mount(r chi.Router) {
	r.Post("/api/user/orders", h.handle)
}

func (h *UploadOrderNumber) handle(rw http.ResponseWriter, req *http.Request) {
	const op = "uploadOrderNumber"

	ctx := req.Context()

	token, ok := getJwtFromContext(ctx, op)
	if !ok {
		writeInternalServerError(rw)

		return
	}

	userID, ok := getUserIDFromJwt(token, op)
	if !ok {
		writeInternalServerError(rw)

		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		writeInternalServerError(rw)

		return
	}

	orderNumber := string(body)

	err = h.usecase.Do(ctx, userID, orderNumber)
	if err != nil {
		{
			var e *usecases.ValidationError

			var e1 *usecases.BusinessError

			switch {
			case errors.As(err, &e):
				writeValidationError(
					rw,
					http.StatusUnprocessableEntity,
					e.Code,
					e.Field,
					e.Message,
				)
			case errors.As(err, &e1):
				var statusCode int

				if e1.IsCodeMatch("D1531") {
					statusCode = http.StatusOK
				}

				if e1.IsCodeMatch("D1426") {
					statusCode = http.StatusConflict
				}

				writeBusinessError(
					rw,
					statusCode,
					e1.Code,
					e1.Message,
				)
			default:
				writeInternalServerError(rw)

				return
			}
		}

		return
	}

	rw.WriteHeader(http.StatusAccepted)
}
