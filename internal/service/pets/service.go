package pets

import (
	"context"
	"fmt"

	"github.com/ahawi/diploma/internal/entity"
)

type petRepository interface {
	PetsBatch(ctx context.Context, lastID int64, offset int64, filter entity.PetFilter) ([]*entity.PetShort, error)
	PetByID(ctx context.Context, id int64) (*entity.Pet, error)
	AddPet(ctx context.Context, pet *entity.Pet) error
}

type Service struct {
	petRepository petRepository
}

func NewService(petRepository petRepository) *Service {
	return &Service{
		petRepository: petRepository,
	}
}

func (s *Service) PetsBatch(ctx context.Context, lastID int64, offset int64, filter entity.PetFilter) ([]*entity.PetShort, error) {
	return s.petRepository.PetsBatch(ctx, lastID, offset, filter)
}

func (s *Service) PetByID(ctx context.Context, id int64) (*entity.Pet, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: invalid petID", entity.ErrBadRequest)
	}
	pet, err := s.petRepository.PetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("petRepository.PetByID: %w", err)
	}

	return pet, nil
}

func (s *Service) AddPet(ctx context.Context, pet *entity.Pet) error {
	switch {
	case len(pet.Name) == 0:
		return fmt.Errorf("%w: invalid name", entity.ErrBadRequest)
	case len(pet.Gender) == 0:
		return fmt.Errorf("%w: invalid gender", entity.ErrBadRequest)
	case len(pet.Breed) == 0:
		return fmt.Errorf("%w: invalid breed", entity.ErrBadRequest)
	case len(pet.Color) == 0:
		return fmt.Errorf("%w: invalid color", entity.ErrBadRequest)
	case len(pet.Personality) == 0:
		return fmt.Errorf("%w: invalid personality", entity.ErrBadRequest)
	case len(pet.Health) == 0:
		return fmt.Errorf("%w: invalid health", entity.ErrBadRequest)
	case len(pet.Experience) == 0:
		return fmt.Errorf("%w: invalid experience", entity.ErrBadRequest)
	case pet.WoolLength <= 0:
		return fmt.Errorf("%w: invalid wool lenght", entity.ErrBadRequest)
	case pet.Size <= 0:
		return fmt.Errorf("%w: invalid size", entity.ErrBadRequest)
	case pet.Age <= 0:
		return fmt.Errorf("%w: invalid age", entity.ErrBadRequest)
	case pet.Weight <= 0:
		return fmt.Errorf("%w: invalid weight", entity.ErrBadRequest)
	}

	err := s.petRepository.AddPet(ctx, pet)
	if err != nil {
		return fmt.Errorf("petRepository.AddPet: %w", err)
	}

	return nil
}
