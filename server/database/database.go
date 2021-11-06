package database

import (
	"database/sql"
	"fmt"
	"server/cats"
	"server/users"
)

type DatabaseInfo struct {
	Host string
	Port int
	User string
	Password string
	Dbname string
}

type DB struct {
	*sql.DB
}

// No pointers for interfaces?
type DogHandler interface {
	Dog(s string) *sql.Row
}

type DogStore struct {
	db *sql.DB
	Query DogHandler
}

type Store struct {
	Dogs DogStore
	Cats cats.CatStore
	Users users.UserStore
}

type StoreHandler interface {
	DogHandler
	cats.CatHandler
	users.UserHandler
}

func (store DogStore) Dog (s string) *sql.Row {
	row := store.db.QueryRow(s)
	return row
}

func OpenDB(info DatabaseInfo) (*Store, error) {
	connInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	  "password=%s dbname=%s sslmode=disable",
	  info.Host, info.Port, info.User, info.Password, info.Dbname)

	db, err := sql.Open("postgres", connInfo)

	if err != nil {
	  return nil, err
	}
  
	err = db.Ping()

	if err != nil {
	  return nil, err
	}

	return &Store{ Dogs: DogStore{db: db}, Cats: cats.CatStore{Db: db}, Users: users.UserStore{Db: db}}, nil
}