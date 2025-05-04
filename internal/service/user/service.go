package user

import (
	"context"
	"fmt"

	"github.com/ahawi/diploma/internal/entity"
)

type userRepository interface {
	AddForm(ctx context.Context, userID int64, form *entity.Form) error
	AddInteraction(ctx context.Context, interaction *entity.Interaction) error
	AddUser(ctx context.Context, name string) error
}

type Service struct {
	userRepository userRepository
}

func NewService(userRepository userRepository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}

func (s *Service) AddForm(ctx context.Context, userID int64, form *entity.Form) error {
	if userID <= 0 {
		return fmt.Errorf("%w: invalid userID", entity.ErrBadRequest)
	}
	err := s.userRepository.AddForm(ctx, userID, form)
	if err != nil {
		return fmt.Errorf("userRepository.AddForm: %w", err)
	}

	return nil
}

func (s *Service) AddInteraction(ctx context.Context, interaction *entity.Interaction) error {
	if interaction.UserID <= 0 {
		return fmt.Errorf("%w: invalid userID", entity.ErrBadRequest)
	}
	if interaction.PetID <= 0 {
		return fmt.Errorf("%w: invalid petID", entity.ErrBadRequest)
	}

	err := s.userRepository.AddInteraction(ctx, interaction)
	if err != nil {
		return fmt.Errorf("userRepository.AddInteraction: %w", err)
	}

	return nil
}

func (s *Service) RegisterUser(ctx context.Context, name string) error {
	if len(name) == 0 {
		return fmt.Errorf("%w: invalid name", entity.ErrBadRequest)
	}

	err := s.userRepository.AddUser(ctx, name)
	if err != nil {
		return fmt.Errorf("userRepository.AddUser: %w", err)
	}

	return nil
}
