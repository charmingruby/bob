package {{ .Module }}

import (
	"database/sql"

	"{{ .SourcePath }}/{{ .Module }}/core/repository"
	"{{ .SourcePath }}/{{ .Module }}/core/service"
	"{{ .SourcePath }}/{{ .Module }}/database/postgres"
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

func New{{ .UpperCaseRepository }}Repository(db *sql.DB) repository.{{ .UpperCaseRepository }}Repository {
	return postgres.New{{ .UpperCaseRepository }}Repository(db)
}

func NewHTTPHandler(r *chi.Mux, service *service.Service) {
	endpoint.New(r, service).Register()
}
