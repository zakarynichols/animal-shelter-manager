package app

import (
	"server/cats"
	"server/donations"
	"server/users"

	"github.com/gorilla/mux"
)

// Router returns a new mux router instance
// pointing to the Server struct
func (s *Server) Router() *mux.Router {
	return s.router
}

// Method points to Server struct to give routes access to dependencies.
func (s *Server) InitializeRoutes() {
	// Authentication
	s.router.HandleFunc("/register", users.Register(s.store.Users)).Methods("POST", "OPTIONS")
	s.router.HandleFunc("/logged-in", users.Login(s.store.Users)).Methods("POST", "OPTIONS")

	// Donations
	s.router.HandleFunc("/donations/donate", donations.DonateHandler()).Methods("POST")
	s.router.HandleFunc("/read-cookie", donations.ReadCookieHandler()).Methods("GET")

	cats.Routes(s.router, s.store.Cats, &cats.CatService{})
}