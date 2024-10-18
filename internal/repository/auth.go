package repository

import (
	"context"
	"database/sql"
	"errors"
	"project/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) (*UserRepository, error) {
	if db == nil {
		return nil, errors.New("missing db")
	}

	return &UserRepository{
		db: db,
	}, nil
}

func (s *UserRepository) FindByID(ctx context.Context, id int) (entity.User, error) {
	return entity.User{}, nil
}
