package user

import (
	"context"
	"database/sql"
	"errors"
	"project/entity"
	"project/internal/repository"
)

type Domain struct {
	db *sql.DB

	authRepository *repository.UserRepository
}

func NewDomain(db *sql.DB, authRepository *repository.UserRepository) (*Domain, error) {
	if db == nil {
		return nil, errors.New("missing db")
	}

	if authRepository == nil {
		return nil, errors.New("missing auth repository")
	}

	return &Domain{
		db: db,
	}, nil
}

func (s *Domain) FindByID(ctx context.Context, id int) (entity.User, error) {
	return s.authRepository.FindByID(ctx, id)
}
