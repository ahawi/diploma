package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/ahawi/diploma/internal/api"
	"github.com/ahawi/diploma/internal/config"
	"github.com/ahawi/diploma/internal/repository"
	"github.com/ahawi/diploma/internal/service/pets"
	"github.com/ahawi/diploma/internal/service/recommendation"
	"github.com/ahawi/diploma/internal/service/user"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.MustLoad()
	log.Printf("config: %v", cfg)

	conn, err := repository.NewConnection(ctx, cfg.PostgresConfig)
	if err != nil {
		log.Printf("failed to connect to postgres: %v", err)
	}

	petRepository := repository.NewPetRepository(conn)
	userRepository := repository.NewUserRepository(conn)

	petsService := pets.NewService(petRepository)
	userService := user.NewService(userRepository)
	recommendationService := recommendation.New(
		cfg.ServerConfig.ContentWeight,
		cfg.ServerConfig.CollabWeight,
		userRepository,
		petRepository,
	)

	api := api.NewAPI(userService, petsService, recommendationService, cfg.ServerConfig.TimeoutAPI)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Route("/pets", func(r chi.Router) {
		r.Get("/simple", api.GetPetsBatchHandler)
		r.Get("/rec", api.GetRecommendationPetsBatchHandler)
		r.Get("/{petID}", api.GetPetByIDHandler)
	})

	router.Route("/user", func(r chi.Router) {
		r.Get("/profile", api.GetProfileHandler)

		r.Post("/form", api.PostFormHandler)
		r.Post("/add-interaction", api.PostAddInteractionHandler)
	})

	log.Print("start listening")
	log.Fatal(http.ListenAndServe("[::]:"+cfg.ServerConfig.Port, router))
}
