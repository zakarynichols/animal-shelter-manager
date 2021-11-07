package users

import (
	"database/sql"
	"errors"
	"time"
)

// No pointers for interfaces?
type UserHandler interface {
	CanCreateUser(s string) error
	User(id int) (*RegisteredUser, error)
	CreateUser(username string, hashedPassword []byte, sessionId string) (*RegisteredUser, error)
}

type UserStore struct {
	Db *sql.DB
}

func (store UserStore) CanCreateUser(s string) error {
	row, err := store.Db.Query("select username from users where username = $1", s)

	if err != nil {
		return err
	}

	defer row.Close()

	row.Next()

	var username sql.NullString

	err = row.Scan(&username)

	var errRowsClosed = errors.New("sql: Rows are closed")

	if err != nil && err.Error() != errRowsClosed.Error() {
		return err
	}

	if username.Valid {
		var errUserExists = errors.New("app: User already exists")
		return errUserExists
	}

	return nil
}

func (store UserStore) User(id int) (*RegisteredUser, error) {
	var err error

	row, err := store.Db.Query("select * from users where user_id = $1", id)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	row.Next()

	newUser := RegisteredUser{}

	err = row.Scan(&newUser.Username, &newUser.CreatedOn, &newUser.LastLogin, &newUser.Session)

	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (store UserStore) CreateUser(username string, hashedPassword []byte, newSessionId string) (*RegisteredUser, error) {
	var err error

	row, err := store.Db.Query("insert into users(username, password, last_login, session) values ($1, $2, $3, $4) returning username, created_on, last_login, session", username, hashedPassword, time.Now(), newSessionId)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	row.Next()

	newUser := RegisteredUser{}

	err = row.Scan(&newUser.Username, &newUser.CreatedOn, &newUser.LastLogin, &newUser.Session)

	if err != nil {
		return nil, err
	}

	return &newUser, nil
}