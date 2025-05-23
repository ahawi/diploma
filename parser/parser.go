package parser

import (
	"encoding/json"
	"log"
	"net/http"
	"parser/db"
	"parser/models"
)

func ParseAndSaveDogsFromAPI(apiURL string) {
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Println("Failed to fetch API:", err)
		return
	}
	defer resp.Body.Close()

	var dogs []models.Dog
	if err := json.NewDecoder(resp.Body).Decode(&dogs); err != nil {
		log.Println("Failed to decode JSON:", err)
		return
	}

	for _, dog := range dogs {
		if err := db.SaveDog(dog); err != nil {
			log.Println("Failed to save dog:", dog.Name, err)
		} else {
			log.Println("Saved dog:", dog.Name)
		}
	}
}
