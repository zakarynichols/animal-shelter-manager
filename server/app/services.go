package app

import (
	"server/cats"
	"server/dogs"
	"server/donations"
	"server/users"
)

// Initialize the app service. This includes setting up routes and injecting
// dependencies as parameters to the routes. This method (InitAppService) points
// to Server struct to give routes access to dependencies.
func (s *Server) InitAppService() {
	// User routes
	users.Routes(s.router, s.store.UserStore, s.handlers.UserHandler)
	// Dogs routes
	dogs.Routes(s.router, s.store.DogStore, s.handlers.DogHandler)
	// Cat routes
	cats.Routes(s.router, s.store.CatStore, s.handlers.CatHandler)
	// Donation routes
	donations.Routes(s.router, s.handlers.DonationHandler)
}
