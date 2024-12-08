package rest

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type RestServer struct {
	HttpServer *http.Server
	Router     *chi.Mux
}

func NewServer(port string, router *chi.Mux) *RestServer {
	httpServer := http.Server{
		Addr: ":" + port,
	}

	attachBaseMiddlewares(router)

	return &RestServer{
		HttpServer: &httpServer,
		Router:     router,
	}
}

func (s *RestServer) Run() error {
	if err := http.ListenAndServe(s.HttpServer.Addr, s.Router); err != nil {
		return err
	}

	return nil
}

func (s *RestServer) Shutdown(ctx context.Context) error {
	return s.HttpServer.Shutdown(ctx)
}
