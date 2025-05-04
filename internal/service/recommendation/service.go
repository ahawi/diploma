package recommendation

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"sort"
	"sync"

	"github.com/ahawi/diploma/internal/entity"
)

type userRepository interface {
	GetInteractionsForUser(ctx context.Context, userID int64) ([]*entity.Interaction, error)
	GetSimilarUsers(ctx context.Context, userID int64, threshold float64) ([]*entity.User, error)
}

type petRepository interface {
	PetByID(ctx context.Context, petID int64) (*entity.Pet, error)
	PetsByIDs(ctx context.Context, petIDs []int64) ([]*entity.Pet, error)
	AllPets(ctx context.Context) ([]*entity.Pet, error)
}

type Service struct {
	contentWeight float64
	collabWeight  float64

	userRepository userRepository
	petRepository  petRepository
}

func New(contentWeight, collabWeight float64, userRepository userRepository, petRepository petRepository) *Service {
	return &Service{
		contentWeight: contentWeight,
		collabWeight:  collabWeight,

		userRepository: userRepository,
		petRepository:  petRepository,
	}
}

func (rs *Service) GetRecommendations(ctx context.Context, userID int64, page int, pageSize int) ([]*entity.PetShort, error) {
	if userID <= 0 {
		return nil, fmt.Errorf("%w: invalid userID", entity.ErrBadRequest)
	}

	interactions, err := rs.userRepository.GetInteractionsForUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("userRepository.GetInteractionsForUser: %w", err)
	}

	var wg sync.WaitGroup
	var contentRecs, collabRecs []int64
	var contentErr, collabErr error

	wg.Add(2)
	go func() {
		defer wg.Done()
		contentRecs, contentErr = rs.contentBasedRecommendations(ctx, interactions)
	}()
	go func() {
		defer wg.Done()
		collabRecs, collabErr = rs.collaborativeRecommendations(ctx, userID, interactions)
	}()
	wg.Wait()

	if contentErr != nil || collabErr != nil {
		return nil, errors.Join(contentErr, collabErr)
	}

	combined := rs.mergeRecommendations(contentRecs, collabRecs)
	res, err := rs.paginateResults(ctx, combined, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("paginateResults: %w", err)
	}

	return res, nil
}

func (rs *Service) contentBasedRecommendations(ctx context.Context, interactions []*entity.Interaction) ([]int64, error) {
	userProfile := make(map[string]float64)
	totalWeight := 0.0

	for _, interaction := range interactions {
		pet, err := rs.petRepository.PetByID(ctx, interaction.PetID)
		if err != nil {
			log.Printf("failed to get pet by id: %v", err)
			continue
		}

		weight := interaction.Weight * entity.InteractionWeight(interaction.Type)
		for feature, value := range pet.ToFeatures().Features {
			userProfile[feature] += value * weight
		}
		totalWeight += weight
	}

	if totalWeight > 0 {
		for feature := range userProfile {
			userProfile[feature] /= totalWeight
		}
	}

	allPets, err := rs.petRepository.AllPets(ctx)
	if err != nil {
		return nil, fmt.Errorf("petRepository.AllPets: %w", err)
	}

	type scoredPet struct {
		petID int64
		score float64
	}

	var scoredPets []scoredPet
	for _, pet := range allPets {
		score := cosineSimilarity(userProfile, pet.ToFeatures().Features)
		scoredPets = append(scoredPets, scoredPet{pet.ID, score})
	}

	sort.Slice(scoredPets, func(i, j int) bool {
		return scoredPets[i].score > scoredPets[j].score
	})

	result := make([]int64, len(scoredPets))
	for i, si := range scoredPets {
		result[i] = si.petID
	}

	return result, nil
}

func (rs *Service) collaborativeRecommendations(ctx context.Context, userID int64, interactions []*entity.Interaction) ([]int64, error) {
	similarUsers, err := rs.userRepository.GetSimilarUsers(ctx, userID, 0.5)
	if err != nil {
		return nil, fmt.Errorf("userRepository.GetSimilarUsers: %w", err)
	}

	petScores := make(map[int64]float64)
	for _, su := range similarUsers {
		for _, interaction := range su.Interactions {
			petScores[interaction.PetID] += interaction.Weight * entity.InteractionWeight(interaction.Type)
		}
	}

	seenPets := make(map[int64]struct{})
	for _, interaction := range interactions {
		seenPets[interaction.PetID] = struct{}{}
	}

	var petIDs []int64
	for petID := range petScores {
		if _, ok := seenPets[petID]; !ok {
			petIDs = append(petIDs, petID)
		}
	}

	sort.Slice(petIDs, func(i, j int) bool {
		return petScores[petIDs[i]] > petScores[petIDs[j]]
	})

	return petIDs, nil
}

func (rs *Service) mergeRecommendations(a, b []int64) []int64 {
	combined := make(map[int64]struct{})
	scores := make(map[int64]float64)

	for i, id := range a {
		score := rs.contentWeight * (1.0 / float64(i+1))
		combined[id] = struct{}{}
		scores[id] += score
	}

	for i, id := range b {
		score := rs.collabWeight * (1.0 / float64(i+1))
		combined[id] = struct{}{}
		scores[id] += score
	}

	result := make([]int64, 0, len(combined))
	for id := range combined {
		result = append(result, id)
	}

	sort.Slice(result, func(i, j int) bool {
		return scores[result[i]] > scores[result[j]]
	})

	return result
}

func (rs *Service) paginateResults(ctx context.Context, petIDs []int64, page, pageSize int) ([]*entity.PetShort, error) {
	start := page * pageSize
	if start > len(petIDs) {
		return nil, nil
	}
	end := start + pageSize
	if end > len(petIDs) {
		end = len(petIDs)
	}
	res, err := rs.petRepository.PetsByIDs(ctx, petIDs[start:end])
	if err != nil {
		return nil, fmt.Errorf("petRepository.PetsByIDs: %w", err)
	}
	return entity.Pets(res).ToShort(), nil
}

func cosineSimilarity(a, b map[string]float64) float64 {
	var (
		dotProduct float64 = 0
		magnitudeA float64 = 0
		magnitudeB float64 = 0
	)

	for k, v1 := range a {
		if v2, exists := b[k]; exists {
			dotProduct += v1 * v2
		}
		magnitudeA += v1 * v1
	}

	for _, v := range b {
		magnitudeB += v * v
	}

	if magnitudeA == 0 || magnitudeB == 0 {
		return 0
	}

	return dotProduct / (math.Sqrt(magnitudeA) * math.Sqrt(magnitudeB))
}
