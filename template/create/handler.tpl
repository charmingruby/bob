package endpoint

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

func (e *Endpoint) {{ .HandlerName }}Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write("hello world")
	}
}
