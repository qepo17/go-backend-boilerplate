package admin

import (
	"net/http"
)

func (h *Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user, err := h.authService.FindByID(ctx, 1)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		if user.Password != "password" {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		w.Write([]byte(user.Email))
	}
}
