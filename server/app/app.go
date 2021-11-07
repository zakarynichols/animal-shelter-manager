package app

import (
	"server/config"
	"server/handlers"
	"server/store"

	"github.com/gorilla/mux"
)

type Server struct {
	store     *store.Store
	router *mux.Router
	env *config.Env
	handlers *handlers.Handlers
}

func NewServer(s *store.Store, r *mux.Router, env *config.Env, h *handlers.Handlers) *Server {
	return &Server{store: s, router: r, env: env, handlers: h}
}

func (s *Server) NewRouter() *mux.Router {
	return s.router
}