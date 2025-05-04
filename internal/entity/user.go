package entity

import "time"

type User struct {
	ID           int64
	Interactions []*Interaction
}

type Interaction struct {
	UserID    int64           `db:"user_id"`
	PetID     int64           `db:"pet_id"`
	Type      InteractionType `db:"type"`
	Weight    float64         `db:"weight"`
	CreatedAt time.Time       `db:"created_at"`
}

type InteractionType int

const (
	View InteractionType = iota
	Purchase
	Rating
)

func InteractionWeight(t InteractionType) float64 {
	switch t {
	case View:
		return 0.3
	case Purchase:
		return 1.0
	case Rating:
		return 0.7
	default:
		return 0.5
	}
}
