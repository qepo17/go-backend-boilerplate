package admin

import (
	"project/user"
)

type Handler struct {
	userService *user.Domain
}

func NewHandler(authService *user.Domain) *Handler {
	return &Handler{}
}
