package app

import (
	"server/cats"
	"server/dogs"
	"server/donations"
	"server/users"
)

type handlers struct {
	cats.CatHandler
	dogs.DogHandler
	users.UserHandler
	donations.DonationHandler
}

func NewHandlers() *handlers {
	return &handlers{cats.NewCatHandler(), dogs.NewDogHandler(), users.NewUserHandler(), donations.NewDonationHandler()}
}
