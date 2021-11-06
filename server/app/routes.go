package app

import (
	"server/auth"
	"server/cats"
	"server/donations"

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
	s.router.HandleFunc("/register", auth.SignUp(&s.store.Users)).Methods("POST", "OPTIONS")
	// s.router.HandleFunc("/logged-in", auth.LoggedInRoute(s.db)).Methods("GET", "OPTIONS")

	// Donations
	s.router.HandleFunc("/donations/donate", donations.DonateHandler()).Methods("POST")
	s.router.HandleFunc("/read-cookie", donations.ReadCookieHandler()).Methods("GET")

	s.router.HandleFunc("/cats", cats.GetCatById(s.store.Cats)).Methods("GET")
}