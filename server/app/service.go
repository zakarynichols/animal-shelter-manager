package app

import (
	"server/cats"
	"server/dogs"
	"server/donations"
	"server/users"
)

// Method points to Server struct to give routes access to dependencies.
func (s *Server) InitAppService() {
	// Authentication
	// s.router.HandleFunc("/register", users.Register(s.store.Users)).Methods("POST", "OPTIONS")
	// s.router.HandleFunc("/logged-in", users.Login(s.store.Users)).Methods("POST", "OPTIONS")

	// Donations
	s.router.HandleFunc("/donations/donate", donations.DonateHandler()).Methods("POST")
	s.router.HandleFunc("/read-cookie", donations.ReadCookieHandler()).Methods("GET")

	users.Routes(s.router, s.store.Users, s.handlers.Users)
	dogs.Routes(s.router, s.store.Dogs, s.handlers.Dogs)
	cats.Routes(s.router, s.store.Cats, s.handlers.Cats)
}
