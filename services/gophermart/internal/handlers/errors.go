package handlers

import (
	"net/http"

	json "github.com/bytedance/sonic"
)

type ValidationErrorResponse struct {
	Code    string `json:"code"`
	Field   string `json:"field"`
	Message string `json:"message"`
}

type BusinessErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type InternalServerErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func writeValidationError(
	rw http.ResponseWriter,
	statusCode int,
	code, field, message string,
) {
	rw.WriteHeader(statusCode)
	enc := json.ConfigDefault.NewEncoder(rw)
	enc.Encode(
		ValidationErrorResponse{
			Code:    code,
			Field:   field,
			Message: message,
		},
	)
}

func writeBusinessError(
	rw http.ResponseWriter,
	statusCode int,
	code, message string,
) {
	rw.WriteHeader(statusCode)
	enc := json.ConfigDefault.NewEncoder(rw)
	enc.Encode(
		BusinessErrorResponse{
			Code:    code,
			Message: message,
		},
	)
}

func writeInternalServerError(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusInternalServerError)
	enc := json.ConfigDefault.NewEncoder(rw)
	enc.Encode(
		InternalServerErrorResponse{
			Code:    "S1394",
			Message: "internal error",
		},
	)
}
