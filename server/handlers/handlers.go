package handlers

import "server/cats"

type Handlers struct {
	Cats cats.CatHandler
}

func NewHandlers() *Handlers {
	return &Handlers{Cats: &cats.CatService{}}
}