package app

import (
	"server/config"
	"server/database"

	"github.com/gorilla/mux"
)

type Server struct {
	store     *database.Store
	router *mux.Router
	env *config.Env
}

func NewServer(s *database.Store, r *mux.Router, env *config.Env) *Server {
	return &Server{store: s, router: r, env: env}
}