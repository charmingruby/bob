package {{ .Module }}

import (
	"{{ .SourcePath }}/{{ .Module }}/core/service"
	"{{ .SourcePath }}/{{ .Module }}/database/{{ .Database }}"
	"{{ .SourcePath }}/{{ .Module }}/transport/rest/endpoint"
	"github.com/go-chi/chi/v5"
)

func NewService() *service.Service {
	return service.New(service.Input{
		{{ .CapitalizedRepositoryName }}Repository: {{ .Database }}.New{{ .CapitalizedRepositoryName }}Repository(),
	})
}

func NewHTTPHandler(r *chi.Mux, service *service.Service) {
	endpoint.New(r, service).Register()
}
