package handlers

import (
	"context"
	"errors"
	"net/http"

	json "github.com/bytedance/sonic"
	"github.com/go-chi/chi/v5"

	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/usecases"
)

type LoginUsecase interface {
	Do(
		ctx context.Context,
		login,
		password string,
	) (token string, err error)
}

type Login struct {
	usecase LoginUsecase
}

func NewLogin(usecase LoginUsecase) *Login {
	return &Login{usecase: usecase}
}

func (h *Login) Mount(r chi.Router) {
	r.Post("/api/user/login", h.handle)
}

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (h *Login) handle(rw http.ResponseWriter, req *http.Request) {
	var reqData LoginRequest

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

	ctx := req.Context()

	token, err := h.usecase.Do(ctx, reqData.Login, reqData.Password)
	if err != nil {
		{
			var e *usecases.ValidationError

			var e1 *usecases.BusinessError

			switch {
			case errors.As(err, &e):
				writeValidationError(
					rw,
					http.StatusBadRequest,
					e.Code,
					e.Field,
					e.Message,
				)
			case errors.As(err, &e1):
				writeBusinessError(
					rw,
					http.StatusUnauthorized,
					e1.Code,
					e1.Message,
				)
			default:
				writeInternalServerError(rw)
			}

			return
		}
	}

	rw.Header().Set("Authorization", "Bearer "+token)
	rw.WriteHeader(http.StatusOK)
}
