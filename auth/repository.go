package auth

import (
	"context"
	"project/entity"
)

func (s *Domain) FindByID(ctx context.Context, id int) (entity.User, error) {
	return entity.User{}, nil
}
