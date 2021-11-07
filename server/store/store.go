package store

// https://golang.org/doc/database/querying

import (
	"database/sql"
	"server/cats"
	"server/dogs"
	"server/users"
)

type Store struct {
	Dogs  dogs.DogQuerier
	Cats  cats.CatQuerier
	Users users.UserQuerier
}

func NewStore(db *sql.DB) *Store {
	return &Store{Dogs: &dogs.DogStore{Db: db}, Cats: &cats.CatStore{Db: db}, Users: &users.UserStore{Db: db}}
}