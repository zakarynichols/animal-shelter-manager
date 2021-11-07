package cats

import (
	"database/sql"
)

// No pointers for interfaces?
type CatStore interface {
	cat(id int) (*cat, error)
}

type catStore struct {
	*sql.DB
}

type cat struct {
	CatId int    `json:"cat_id"`
	Name  string `json:"name"`
}

// Instantiate a new store to keep
// db a private field on the struct
func NewCatStore(db *sql.DB) *catStore {
	return &catStore{db}
}

func (store *catStore) cat(id int) (*cat, error) {
	var err error

	row, err := store.Query("select * from cats where cat_id = $1", id)

	row.Next()

	if err != nil {
		return nil, err
	}

	c := cat{}

	err = row.Scan(&c.CatId, &c.Name)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
