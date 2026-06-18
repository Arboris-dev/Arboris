package api

import (
	"Arboris/go_server/config"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func NewHookRouter(envVar config.Config) (*chi.Mux, error) {
	hookRouter := chi.NewRouter()

	hookRouter.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: false,
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
	}))

	hookRouter.Use()

	return hookRouter, nil
}
