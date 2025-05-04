package repository

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ahawi/diploma/internal/entity"
)

type PetRepository struct {
	repository
}

func NewPetRepository(conn *pgxpool.Pool) *PetRepository {
	return &PetRepository{
		repository{
			conn:    conn,
			builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		},
	}
}

func (r *PetRepository) PetsBatch(ctx context.Context, lastID int64, offset int64) ([]*entity.PetShort, error) {
	query := r.builder.Select("id, name, breed, gender, age").
		Where(squirrel.GtOrEq{"id": lastID}).
		OrderBy("id ASC").
		Limit(uint64(offset)).
		Offset(uint64(offset))

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("get sql query: %w", err)
	}

	var result []*entity.PetShort
	err = pgxscan.Select(ctx, r.conn, &result, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("get from db: %w", err)
	}

	return result, nil
}

func (r *PetRepository) PetByID(ctx context.Context, id int64) (*entity.Pet, error) {
	query := r.builder.Select("*").
		Where(squirrel.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("get sql query: %w", err)
	}

	var result []*entity.Pet
	err = pgxscan.Select(ctx, r.conn, &result, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("get from db: %w", err)
	}
	if len(result) == 0 {
		return nil, ErrNotFound
	}

	return result[0], nil
}

func (r *PetRepository) PetsByIDs(ctx context.Context, petIDs []int64) ([]*entity.Pet, error) {
	query := r.builder.Select("*").
		Where(squirrel.Eq{"id": petIDs})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("get sql query: %w", err)
	}

	var result []*entity.Pet
	err = pgxscan.Select(ctx, r.conn, &result, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("get from db: %w", err)
	}

	return result, nil
}

func (r *PetRepository) AllPets(ctx context.Context) ([]*entity.Pet, error) {
	query := r.builder.Select("*")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("get sql query: %w", err)
	}

	var result []*entity.Pet
	err = pgxscan.Select(ctx, r.conn, &result, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("get from db: %w", err)
	}

	return result, nil
}
