package rest

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    int         `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func newResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	w.Header().Add("Content-Type", "application/json")

	var res Response
	res.Message = message

	if data != nil {
		res.Data = data
	}

	jsonRes, _ := json.Marshal(res)

	w.WriteHeader(code)
	w.Write(jsonRes)
}

func OKResponse(w http.ResponseWriter, message string, data interface{}) {
	newResponse(w, http.StatusOK, message, data)
}

func NotFoundErrorResponse(w http.ResponseWriter, message string) {
	newResponse(w, http.StatusNotFound, message, nil)
}

func InternalServerErrorResponse(w http.ResponseWriter) {
	newResponse(w, http.StatusInternalServerError, "internal server error", nil)
}

func BadRequestErrorResponse(w http.ResponseWriter, message string) {
	newResponse(w, http.StatusBadRequest, message, nil)
}
