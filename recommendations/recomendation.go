package service

import (
	"recommendation/data"
	"recommendation/models"
)

func RecommendDogs(userID int) models.Recommendation {
	var user *models.User
	for _, u := range data.Users {
		if u.ID == userID {
			user = &u
			break
		}
	}
	if user == nil {
		return models.Recommendation{UserID: userID, DogIDs: []int{}}
	}

	var recommended []int
	for _, dog := range data.Dogs {
		score := 0

		// Контентная фильтрация
		for _, breed := range user.PreferredBreeds {
			if dog.Breed == breed {
				score += 3
				break
			}
		}
		if dog.Size == user.PreferredSize {
			score += 2
		}
		if dog.Age >= user.AgeRange[0] && dog.Age <= user.AgeRange[1] {
			score += 2
		}
		if dog.Activity == user.ActivityLevel {
			score += 1
		}
		if dog.Temperament == user.Temperament {
			score += 1
		}

		// Пороговое значение
		if score >= 5 {
			recommended = append(recommended, dog.ID)
		}
	}

	return models.Recommendation{UserID: userID, DogIDs: recommended}
}
