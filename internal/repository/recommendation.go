package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RecommendationRepository struct {
	repository
}

func NewRecommendationRepository(conn *pgxpool.Pool) *RecommendationRepository {
	return &RecommendationRepository{
		repository{
			conn:    conn,
			builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		},
	}
}
