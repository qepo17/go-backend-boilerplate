package admin

import "project/auth"

type Handler struct {
	authService *auth.Domain
}

func NewHandler(authService *auth.Domain) *Handler {
	return &Handler{}
}
