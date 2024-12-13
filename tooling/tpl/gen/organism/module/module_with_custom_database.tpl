package {{ .Module }}

import (
	"{{ .SourcePath }}/{{ .Module }}/core/repository"
	"{{ .SourcePath }}/{{ .Module }}/core/service"
	"{{ .SourcePath }}/{{ .Module }}/database/{{ .Database }}"
	"{{ .SourcePath }}/{{ .Module }}/transport/rest/endpoint"
	"github.com/go-chi/chi/v5"
)

func NewService(
	{{ .LowerCaseRepository }}Repository repository.{{ .UpperCaseRepository }}Repository,
) *service.Service {
	return service.New(
		{{ .LowerCaseRepository }}Repository,
	)
}

func New{{ .UpperCaseRepository }}Repository() repository.{{ .UpperCaseRepository }}Repository {
	return mysql.New{{ .UpperCaseRepository }}Repository()
}

func NewHTTPHandler(r *chi.Mux, service *service.Service) {
	endpoint.New(r, service).Register()
}
