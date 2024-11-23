package {{ .Package }}

import (
	"net/http"
)

func ({{ .PackageRegistryIdentifier }} *{{ .PackageRegistry }}) make{{ .Name }}() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello world"))
	}
}
