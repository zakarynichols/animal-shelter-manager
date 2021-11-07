package cats

import (
	"database/sql"
)

// No pointers for interfaces?
type CatHandler interface {
	Cat(id int) (*Cat, error)
}

type CatStore struct {
	Db *sql.DB
}

func (store CatStore) Cat(id int) (*Cat, error) {
	var err error

	row, err := store.Db.Query("select * from cats where cat_id = $1", id)

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