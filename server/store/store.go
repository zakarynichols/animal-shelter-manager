package store

// https://golang.org/doc/database/querying

import (
	"database/sql"
	"server/cats"
	"server/dogs"
	"server/users"
)

type Store struct {
	dogs.DogStore
	cats.CatStore
	users.UserStore
}

func NewStore(db *sql.DB) *Store {
	return &Store{dogs.NewDogStore(db), cats.NewCatStore(db), users.NewUserStore(db)}
}
