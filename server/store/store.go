package store

// https://golang.org/doc/database/querying

import (
	"database/sql"
	"server/cats"
	"server/dogs"
	"server/users"
)

type Store struct {
	Dogs  dogs.DogStore
	Cats  cats.CatStore
	Users users.UserStore
}

func NewStore(db *sql.DB) *Store {
	return &Store{Dogs: dogs.NewDogStore(db), Cats: cats.NewCatStore(db), Users: users.NewUserStore(db)}
}
