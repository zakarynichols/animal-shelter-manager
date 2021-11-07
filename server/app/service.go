package app

import (
	"server/cats"
	"server/donations"
	"server/users"
)

// Method points to Server struct to give routes access to dependencies.
func (s *Server) InitAppService() {
	// Authentication
	s.router.HandleFunc("/register", users.Register(s.store.Users)).Methods("POST", "OPTIONS")
	s.router.HandleFunc("/logged-in", users.Login(s.store.Users)).Methods("POST", "OPTIONS")

	// Donations
	s.router.HandleFunc("/donations/donate", donations.DonateHandler()).Methods("POST")
	s.router.HandleFunc("/read-cookie", donations.ReadCookieHandler()).Methods("GET")

	cats.Routes(s.router, s.store.Cats, s.handlers.Cats)
}