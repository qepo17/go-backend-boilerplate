package config

import (
	"database/sql"
)

func (cfg *Config) NewDB() (*sql.DB, error) {
	// Connect to the database here
	return &sql.DB{}, nil
}
