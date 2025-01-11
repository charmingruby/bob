package endpoint

import (
	"github.com/go-chi/chi/v5"
	"{{ .SourcePath }}/{{ .Module }}/core/service"
)

type Endpoint struct {
	router  *chi.Mux
	service *service.Service
}

func New(r *chi.Mux, service *service.Service) *Endpoint {
	return &Endpoint{
		router:  r,
		service: service,
	}
}

func (e *Endpoint) Register() {
	e.router.Post("/{{ .Module }}/greeting", e.makeGreetingHandler())
}
