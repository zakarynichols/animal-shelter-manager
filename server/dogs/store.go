package dogs

import (
	"database/sql"
)

// No pointers for interfaces?
type DogQuerier interface {
	dog(id int) (*Dog, error)
}

type DogStore struct {
	*sql.DB
}

func NewDogStore(db *sql.DB) *DogStore {
	return &DogStore{db}
}

func (store *DogStore) dog(id int) (*Dog, error) {
	var err error

	row, err := store.Query("select * from dogs where dog_id = $1", id)

	row.Next()

	if err != nil {
		return nil, err
	}

	c := Dog{}

	err = row.Scan(&c.DogId, &c.Name)

	if err != nil {
		return nil, err
	}

	return &c, nil
}