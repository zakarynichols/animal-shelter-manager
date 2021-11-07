package store

// https://golang.org/doc/database/querying

import (
	"database/sql"
	"server/cats"
	"server/users"
)

type DogHandler interface {
	Dog(s string) *sql.Row
}

type DogStore struct {
	db    *sql.DB
	Query DogHandler
}

type Store struct {
	Dogs  DogStore
	Cats  cats.CatHandler
	Users users.UserHandler
}

type StoreHandler interface {
	DogHandler
	cats.CatHandler
	users.UserHandler
}

func (store DogStore) Dog(s string) *sql.Row {
	row := store.db.QueryRow(s)
	return row
}

func NewStore(db *sql.DB) *Store {
	return &Store{Dogs: DogStore{db: db}, Cats: cats.CatStore{Db: db}, Users: users.UserStore{Db: db}}
}