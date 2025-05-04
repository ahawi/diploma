package api

import (
	"context"
	"time"

	"github.com/ahawi/diploma/internal/entity"
)

type userService interface {
	AddForm(ctx context.Context, userID int64, form *entity.Form) error
	AddInteraction(ctx context.Context, interaction *entity.Interaction) error
}

type petService interface {
	PetsBatch(ctx context.Context, lastID int64, offset int64) ([]*entity.PetShort, error)
	PetByID(ctx context.Context, id int64) (*entity.Pet, error)
}

type recommendationService interface {
	GetRecommendations(ctx context.Context, userID int64, page int, pageSize int) ([]*entity.PetShort, error)
}

type API struct {
	userService           userService
	petService            petService
	recommendationService recommendationService

	timeout time.Duration
}

func NewAPI(
	userService userService,
	petService petService,
	recommendationService recommendationService,
	timeout time.Duration,
) *API {
	return &API{
		userService:           userService,
		petService:            petService,
		recommendationService: recommendationService,

		timeout: timeout,
	}
}
