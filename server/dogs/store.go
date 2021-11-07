package dogs

import (
	"database/sql"
)

// No pointers for interfaces?
type DogStore interface {
	dog(id int) (*dog, error)
}

type dogStore struct {
	*sql.DB
}

type dog struct {
	DogId int    `json:"dog_id"`
	Name  string `json:"name"`
}

func NewDogStore(db *sql.DB) *dogStore {
	return &dogStore{db}
}

func (store *dogStore) dog(id int) (*dog, error) {
	var err error

	row, err := store.Query("select * from dogs where dog_id = $1", id)

	row.Next()

	if err != nil {
		return nil, err
	}

	c := dog{}

	err = row.Scan(&c.DogId, &c.Name)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
