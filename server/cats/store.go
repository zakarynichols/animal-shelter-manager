package cats

import (
	"database/sql"
)

// No pointers for interfaces?
type CatQuerier interface {
	cat(id int) (*Cat, error)
}

type CatStore struct {
	*sql.DB
}

// Instantiate a new store to keep
// db a private field on the struct
func NewCatStore(db *sql.DB) *CatStore {
	return &CatStore{db}
}

func (store *CatStore) cat(id int) (*Cat, error) {
	var err error

	row, err := store.Query("select * from cats where cat_id = $1", id)

	row.Next()

	if err != nil {
		return nil, err
	}

	c := Cat{}

	err = row.Scan(&c.CatId, &c.Name)

	if err != nil {
		return nil, err
	}

	return &c, nil
}