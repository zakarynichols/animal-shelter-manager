package users

import (
	"database/sql"
	"errors"
	"time"
)

// No pointers for interfaces?

// Expose the interface and and the
// func to instantiate a user store
type UserStore interface {
	canCreateUser(s string) error
	user(id int) (*RegisteredUser, error)
	createUser(username string, hashedPassword []byte, sessionId string) (*RegisteredUser, error)
}

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

type userStore struct {
	*sql.DB
}

func NewUserStore(db *sql.DB) *userStore {
	return &userStore{db}
}

func (store *userStore) canCreateUser(s string) error {
	row, err := store.Query("select username from users where username = $1", s)

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

func (store *userStore) user(id int) (*RegisteredUser, error) {
	var err error

	row, err := store.Query("select * from users where user_id = $1", id)

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

func (store *userStore) createUser(username string, hashedPassword []byte, newSessionId string) (*RegisteredUser, error) {
	var err error

	row, err := store.Query("insert into users(username, password, last_login, session) values ($1, $2, $3, $4) returning username, created_on, last_login, session", username, hashedPassword, time.Now(), newSessionId)

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
