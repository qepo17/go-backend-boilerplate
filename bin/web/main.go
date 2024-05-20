package main

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"project/auth"
	"project/config"
	"project/http/handler/admin"
	"project/http/handler/user"
	"project/http/middleware"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get config")
	}

	db, err := cfg.NewDB()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}
	defer db.Close()

	r := chi.NewRouter()

	r.Use(middleware.Cors().Handler)

	r.Get("/_health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	authService, err := auth.NewDomain(db)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create auth service")
	}

	adminHandler := admin.NewHandler(authService)
	userHandler := user.NewHandler()

	adminHandler.RegisterRoutes(r)
	userHandler.RegisterRoutes(r)

	srv := &http.Server{
		Addr:    net.JoinHostPort("", cfg.Port),
		Handler: r,
	}

	// Listen for OS interrupt signal
	exitSig := make(chan os.Signal, 1)
	signal.Notify(exitSig, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-exitSig
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Error().Err(err).Msg("failed to shutdown server")
		}
	}()

	err = srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error().Err(err).Msg("serving http server")
	}
}
