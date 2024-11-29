package endpoint

import (
	"{{ .SourcePath }}/{{ .Module }}/core/service"
	"github.com/go-chi/chi/v5"
)

type Endpoint struct {
	router *chi.Mux
	service *service.Service
}

func New(r *chi.Mux, service *service.Service) *Endpoint {
	return &Endpoint{
		router: r,
		service: service,
	}
}

func (e *Endpoint) Register() {
	e.router.Get("/{{ .Module }}/ping", e.makePingHandler())
}
