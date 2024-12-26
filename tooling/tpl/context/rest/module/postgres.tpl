package {{ .Module }}

import (
	"{{ .SourcePath }}/{{ .Module }}/core/service"
	"{{ .SourcePath }}/{{ .Module }}/database/postgres"
	"{{ .SourcePath }}/{{ .Module }}/transport/rest/endpoint"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func NewService(db *sqlx.DB) (*service.Service, error) {
	{{ .LowerCaseRepository }}Repo, err := postgres.New{{ .UpperCaseRepository }}Repository(db)
	if err != nil {
		return nil, err
	}
	
	return service.New(service.Input{
		{{ .UpperCaseRepository }}Repository: {{ .LowerCaseRepository }}Repo,
	}), nil	
}

func NewHTTPHandler(r *chi.Mux, service *service.Service) {
	endpoint.New(r, service).Register()
}
