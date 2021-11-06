package users

import (
	"database/sql"
	"time"
)

type PreAuthenticatedUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisteredUser struct {
	Username  string         `json:"username"`
	CreatedOn time.Time      `json:"created_on"`
	LastLogin time.Time      `json:"last_login"`
	Session   sql.NullString `json:"session"`
}
