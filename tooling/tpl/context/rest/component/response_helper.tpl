package rest

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

func newResponse[T any](w http.ResponseWriter, code int, message string, data T) {
	w.Header().Add("Content-Type", "application/json")

	res := Response[T]{
		Message: message,
		Data:    data,
	}

	jsonRes, _ := json.Marshal(res)

	w.WriteHeader(code)
	w.Write(jsonRes)
}

func OKResponse[T any](w http.ResponseWriter, message string, data T) {
	newResponse[*T](w, http.StatusOK, message, &data)
}

func NotFoundErrorResponse[T any](w http.ResponseWriter, message string) {
	newResponse[*T](w, http.StatusNotFound, message, nil)
}

func InternalServerErrorResponse[T any](w http.ResponseWriter) {
	newResponse[*T](w, http.StatusInternalServerError, "internal server error", nil)
}

func BadRequestErrorResponse[T any](w http.ResponseWriter, message string) {
	newResponse[*T](w, http.StatusBadRequest, message, nil)
}
