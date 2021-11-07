package handlers

import (
	"server/cats"
	"server/dogs"
	"server/users"
)

type Handlers struct {
	Cats  cats.CatHandler
	Dogs  dogs.DogHandler
	Users users.UserHandler
}

func NewHandlers() *Handlers {
	return &Handlers{Cats: cats.NewCatHandler(), Dogs: dogs.NewDogHandler(), Users: users.NewUserHandler()}
}
