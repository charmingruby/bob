package {{ .Module }}

import (
	"{{ .SourcePath }}/{{ .Module }}/core/service"
	"{{ .SourcePath }}/{{ .Module }}/database/postgres"
	"{{ .SourcePath }}/{{ .Module }}/transport/rest/endpoint"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func NewService(db *sqlx.DB) (*service.Service, error) {
	{{ .LowerCaseRepositoryName }}Repo, err := postgres.New{{ .CapitalizedRepositoryName }}Repository(db)
	if err != nil {
		return nil, err
	}
	
	return service.New(service.Input{
		{{ .CapitalizedRepositoryName }}Repository: {{ .LowerCaseRepositoryName }}Repo,
	}), nil	
}

func NewHTTPHandler(r *chi.Mux, service *service.Service) {
	endpoint.New(r, service).Register()
}
