package user

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) RegisterRoutes(r *chi.Mux) http.Handler {
	r.Group(func(r chi.Router) {
		r.Post("/user/login", h.Login())
	})

	return r
}
