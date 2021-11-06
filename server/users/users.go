package users

import (
	"database/sql"
)

type DB struct {
	*sql.DB
}