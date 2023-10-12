package apisever

import (
	"anime/internal/config"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type Server struct {
	cfg    *config.Config
	router *chi.Mux
}

func New(cfg *config.Config, router *chi.Mux) *Server {
	return &Server{
		cfg:    cfg,
		router: router,
	}
}

func (s *Server) Run() error {
	log.Default().Printf("server start addresses: %s", "http://localhost:"+s.cfg.ServerPort)
	if err := http.ListenAndServe(":"+s.cfg.ServerPort, s.router); err != nil {
		return err
	}

	return nil
}
