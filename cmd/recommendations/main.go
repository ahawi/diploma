package main

import (
	"context"
	"log"

	"github.com/ahawi/diploma/internal/config"
	"github.com/ahawi/diploma/internal/repository"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.MustLoad()

	conn, err := repository.NewConnection(ctx, cfg.PostgresConfig)
	if err != nil {
		log.Fatal(err)
	}

	_ = conn
}
