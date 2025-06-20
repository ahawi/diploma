package parser

import (
	"database/sql"
	"log"

	"github.com/ahawi/diploma/internal/entity"
	"github.com/ahawi/diploma/internal/repository"
	"github.com/ahawi/diploma/internal/service/parser/shelters"
)

func RunAllParsers(db *sql.DB) error {
	sheltersList, err := repository.GetAllShelters(db)
	if err != nil {
		return err
	}

	for _, sh := range sheltersList {
		log.Printf("Parsing shelter: %s", sh.URL)

		var dogs []entity.Pet
		switch sh.ID {
		case 1:
			dogs, err = shelters.ParseShelter1(sh)
		case 2:
			dogs, err = shelters.ParseShelter2(sh)
		default:
			dogs, err = shelters.ParseGeneric(sh)
		}

		if err != nil {
			log.Printf("Failed to parse %s: %v", sh.URL, err)
			continue
		}

		for _, d := range dogs {
			d.ShelterID = sh.ID
			if err := repository.AddPet(db, d); err != nil {
				log.Printf("Failed to save dog %s: %v", d.Name, err)
			}
		}
	}
	return nil
}
