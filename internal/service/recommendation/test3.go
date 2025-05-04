package recommendation

import (
	"fmt"
	"math"
	"sort"

	"github.com/ahawi/diploma/internal/entity"
)

type recommendationRepository interface {
}

type Test3S struct {
	recommendationRepository recommendationRepository
}

func NewTest3S(recommendationRepository recommendationRepository) *Test3S {
	return &Test3S{
		recommendationRepository: recommendationRepository,
	}
}

type userMatch struct {
	UserID int64
	Rate   float64
}

func dotProduct(vecA, vecB map[int64]float64) float64 {
	d := 0.0
	for dim, valA := range vecA {
		if valB, ok := vecB[dim]; ok {
			d += valA * valB
		}
	}
	return d
}

func distCosine(vecA, vecB map[int64]float64) float64 {
	dot := dotProduct(vecA, vecB)
	magA := math.Sqrt(dotProduct(vecA, vecA))
	magB := math.Sqrt(dotProduct(vecB, vecB))

	if magA == 0 || magB == 0 {
		return 0.0
	}
	return dot / (magA * magB)
}

func (s *Test3S) MostRelevantPets(userID int64, userRates map[int64]map[int64]float64, nBestUsers, nBestProducts int) []entity.Feedback {
	var matches []userMatch
	currentUserRates := userRates[userID]

	for u, rates := range userRates {
		if u != userID {
			rate := distCosine(currentUserRates, rates)
			matches = append(matches, userMatch{u, rate})
		}
	}

	sort.Slice(matches, func(i, j int) bool {
		if matches[i].Rate == matches[j].Rate {
			return matches[i].UserID > matches[j].UserID
		}
		return matches[i].Rate > matches[j].Rate
	})

	if len(matches) > nBestUsers {
		matches = matches[:nBestUsers]
	}

	fmt.Printf("Most correlated with '%d' users:\n", userID)
	for _, match := range matches {
		fmt.Printf("  UserID: %d  Rate: %6.4f\n", match.UserID, match.Rate)
	}

	filteredMatches := make(map[int64]float64)
	simAll := 0.0
	for _, match := range matches {
		if match.Rate > 0 {
			filteredMatches[match.UserID] = match.Rate
			simAll += match.Rate
		}
	}

	sim := make(map[int64]float64)
	for user, rate := range filteredMatches {
		for product, rating := range userRates[user] {
			if _, exists := currentUserRates[product]; !exists {
				sim[product] += rating * rate
			}
		}
	}

	if simAll != 0 {
		for product := range sim {
			sim[product] /= simAll
		}
	} else {
		sim = make(map[int64]float64)
	}

	var products []entity.Feedback
	for product, rate := range sim {
		products = append(products, entity.Feedback{PetID: product, Score: rate})
	}

	sort.Slice(products, func(i, j int) bool {
		if products[i].Score == products[j].Score {
			return products[i].PetID > products[j].PetID
		}
		return products[i].Score > products[j].Score
	})

	if len(products) > nBestProducts {
		products = products[:nBestProducts]
	}

	fmt.Println("Most correlated products:")
	for _, prod := range products {
		fmt.Printf("  ProductID: %d  CorrelationCoeff: %6.4f\n", prod.PetID, prod.Score)
	}

	return products
}
