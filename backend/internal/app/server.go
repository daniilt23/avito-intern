package app

import (
	"avito/internal/config"
	"time"

	"net/http"
)

type Server struct {
	srv *http.Server
}

func NewServer(cfg config.Config, handler http.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Addr:              cfg.Server.Host + ":" + cfg.Server.Port,
			Handler:           handler,
			ReadHeaderTimeout: 10 * time.Second,
		},
	}
}

func (s *Server) Run() error {
	return s.srv.ListenAndServe()
}
