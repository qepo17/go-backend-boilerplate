package auth

import (
	"database/sql"
	"errors"
)

type Domain struct {
	db *sql.DB
}

func NewDomain(db *sql.DB) (*Domain, error) {
	if db == nil {
		return &Domain{}, errors.New("missing db")
	}

	return &Domain{
		db: db,
	}, nil
}
