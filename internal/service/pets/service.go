package pets

import (
	"context"
	"fmt"

	"github.com/ahawi/diploma/internal/entity"
)

type petRepository interface {
	PetsBatch(ctx context.Context, lastID int64, offset int64) ([]*entity.PetShort, error)
	PetByID(ctx context.Context, id int64) (*entity.Pet, error)
}

type Service struct {
	petRepository petRepository
}

func NewService(petRepository petRepository) *Service {
	return &Service{
		petRepository: petRepository,
	}
}

func (s *Service) PetsBatch(ctx context.Context, lastID int64, offset int64) ([]*entity.PetShort, error) {
	return s.petRepository.PetsBatch(ctx, lastID, offset)
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
