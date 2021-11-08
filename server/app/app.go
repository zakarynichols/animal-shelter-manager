package app

import (
	"server/config"
	"server/store"

	"github.com/gorilla/mux"
)

type Server struct {
	store    *store.Store
	router   *mux.Router
	env      *config.Env
	handlers *handlers
}

func NewServer(s *store.Store, r *mux.Router, env *config.Env, h *handlers) *Server {
	return &Server{store: s, router: r, env: env, handlers: h}
}
