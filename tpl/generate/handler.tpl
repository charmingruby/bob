package endpoint

import (
	"net/http"
)

func (e *Endpoint) {{ .HandlerName }}() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello world"))
	}
}
