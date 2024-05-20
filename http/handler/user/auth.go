package user

import "net/http"

func (h *Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Login User"))
	}
}
