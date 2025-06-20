package repository

import (
	"database/sql"

	"github.com/ahawi/diploma/internal/entity"
)

func (r *PetRepository) GetAllShelters(db *sql.DB) ([]entity.Shelter, error) {
	rows, err := db.Query("SELECT id, url, name FROM shelters")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shelters []entity.Shelter
	for rows.Next() {
		var s entity.Shelter
		if err := rows.Scan(&s.ID, &s.URL, &s.Name); err != nil {
			return nil, err
		}
		shelters = append(shelters, s)
	}
	return shelters, nil
}
