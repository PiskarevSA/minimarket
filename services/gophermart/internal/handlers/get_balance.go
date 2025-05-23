package handlers

import (
	"context"
	"net/http"

	json "github.com/bytedance/sonic"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
)

type GetBalanceUsecase interface {
	Do(
		ctx context.Context,
		userId uuid.UUID,
	) (current, withdrawn objects.Amount, err error)
}

type GetBalance struct {
	usecase GetBalanceUsecase
}

func NewGetBalance(usecase GetBalanceUsecase) *GetBalance {
	return &GetBalance{usecase: usecase}
}

func (h *GetBalance) Mount(r chi.Router) {
	r.Get("/api/user/balance", h.handle)
}

type GetBalanceResponse struct {
	Current   string `json:"current"`
	Withdrawn string `json:"withdrawn"`
}

func (h *GetBalance) handle(rw http.ResponseWriter, req *http.Request) {
	const op = "getBalance"

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

	current, withdrawn, err := h.usecase.Do(ctx, userId)
	if err != nil {
		writeInternalServerError(rw)

		return
	}

	respData := GetBalanceResponse{
		Current:   current.String(),
		Withdrawn: withdrawn.String(),
	}

	rw.WriteHeader(http.StatusOK)
	enc := json.ConfigDefault.NewEncoder(rw)
	enc.Encode(respData)
}
