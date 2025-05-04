package repository

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ahawi/diploma/internal/entity"
)

type UserRepository struct {
	repository
}

func NewUserRepository(conn *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		repository{
			conn:    conn,
			builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		},
	}
}

func (r *UserRepository) AddForm(ctx context.Context, userID int64, form *entity.Form) error {
	tx, err := r.conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	for _, fb := range form.Feedbacks {
		query := r.builder.Insert("interactions").
			Columns("user_id", "pet_id", "type", "weight").
			Values(userID, fb.PetID, entity.Rating, fb.Score)

		sql, args, err := query.ToSql()
		if err != nil {
			return fmt.Errorf("get sql query for userID: %d, petID: %d, err: %w", userID, fb.PetID, err)
		}

		_, err = tx.Exec(ctx, sql, args...)
		if err != nil {
			return fmt.Errorf("save feedback for userID: %d, petID: %d, err: %w", userID, fb.PetID, err)
		}
	}

	return tx.Commit(ctx)
}

func (r *UserRepository) GetInteraction(ctx context.Context, userID, petID int64) (*entity.Interaction, error) {
	query := r.builder.Select("user_id, pet_id, type, weight").
		Where(squirrel.Eq{"user_id": userID}).
		Where(squirrel.Eq{"pet_id": petID})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("get sql query: %w", err)
	}

	var res []*entity.Interaction
	err = pgxscan.Select(ctx, r.conn, &res, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("get from db: %w", err)
	}
	if len(res) == 0 {
		return nil, ErrNotFound
	}

	return res[0], nil
}

func (r *UserRepository) GetUser(ctx context.Context, userID int64) (*entity.User, error) {
	interactions, err := r.GetInteractionsForUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("get interactions for user: %w", err)
	}

	return &entity.User{
		ID:           userID,
		Interactions: interactions,
	}, nil
}

func (r *UserRepository) GetInteractionsForUser(ctx context.Context, userID int64) ([]*entity.Interaction, error) {
	query := r.builder.Select("user_id, pet_id, type, weight").
		Where(squirrel.Eq{"user_id": userID})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("get sql query: %w", err)
	}

	var res []*entity.Interaction
	err = pgxscan.Select(ctx, r.conn, &res, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("get from db: %w", err)
	}

	return res, nil
}

func (r *UserRepository) GetSimilarUsers(ctx context.Context, userID int64, threshold float64) ([]*entity.User, error) {
	return nil, nil
}

func (r *UserRepository) AddInteraction(ctx context.Context, interaction *entity.Interaction) error {
	query := r.builder.Insert("interactions").
		Columns("user_id", "pet_id", "type", "weight").
		Values(interaction.UserID, interaction.PetID, interaction.Type, interaction.Weight)

	sql, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("get sql query: %w", err)
	}

	_, err = r.conn.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("save interaction: %w", err)
	}

	return nil
}
