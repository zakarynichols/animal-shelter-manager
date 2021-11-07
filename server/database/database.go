package database

import (
	"database/sql"
	"fmt"
)

type DatabaseInfo struct {
	Host string
	Port int
	User string
	Password string
	Dbname string
}

type DB struct {
	*sql.DB
}

func OpenDB(info DatabaseInfo) (*sql.DB, error) {
	connInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	  "password=%s dbname=%s sslmode=disable",
	  info.Host, info.Port, info.User, info.Password, info.Dbname)

	db, err := sql.Open("postgres", connInfo)

	if err != nil {
	  return nil, err
	}
  
	err = db.Ping()

	if err != nil {
	  return nil, err
	}

	return db, nil
}