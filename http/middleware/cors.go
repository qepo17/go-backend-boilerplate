package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func Cors() *cors.Cors {
	corsAllowedOrigins := []string{
		"http://localhost:3000",
		"http://localhost:8080",
	}

	return cors.New(cors.Options{
		AllowedOrigins:   corsAllowedOrigins,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowedHeaders:   []string{"Authorization"},
		AllowCredentials: true,
		MaxAge:           3600, // 1 hour
	})
}
