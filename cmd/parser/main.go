package main

import (
	"log"
	"os"

	"github.com/robfig/cron/v3"
	"parser-service/parser"
	"parser-service/repository"
)

func main() {
	connStr := os.Getenv("POSTGRES_CONN")
	db, err := repository.ConnectPostgres(connStr)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	defer db.Close()

	c := cron.New()
	_, err = c.AddFunc("@weekly", func() {
		log.Println("Starting scheduled parsing job...")
		if err := parser.RunAllParsers(db); err != nil {
			log.Printf("Parsing job failed: %v", err)
		}
	})
	if err != nil {
		log.Fatalf("Failed to schedule job: %v", err)
	}

	log.Println("Cron started. Waiting for scheduled tasks...")
	c.Start()
	select {} 
}
