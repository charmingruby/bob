package endpoint

import (
	"net/http"
)

func (e *Endpoint) make{{ .Name }}Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello world"))
	}
}
