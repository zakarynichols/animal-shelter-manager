package dogs

import (
	"database/sql"
)

// No pointers for interfaces?
type DogQuerier interface {
	Dog(id int) (*Dog, error)
}

type DogStore struct {
	Db *sql.DB
}

func (store *DogStore) Dog(id int) (*Dog, error) {
	var err error

	row, err := store.Db.Query("select * from dogs where dog_id = $1", id)

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