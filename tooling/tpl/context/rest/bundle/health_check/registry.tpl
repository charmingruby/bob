package endpoint

import "github.com/go-chi/chi/v5"

type Endpoint struct {
	router *chi.Mux
}

func New(r *chi.Mux) *Endpoint {
	return &Endpoint{
		router: r,
	}
}

func (e *Endpoint) Register() {
	e.router.Get("/health-check", e.makeHealthCheckHandler())
}
