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

func (r *PetRepository) PetsBatch(ctx context.Context, lastID int64, offset int64, filter entity.PetFilter) ([]*entity.PetShort, error) {
	query := r.builder.Select("id, name, breed, gender, age").
		From("dogs").
		Where(squirrel.GtOrEq{"id": lastID})

	if filter.Age != nil {
		query = query.Where(squirrel.GtOrEq{"age": *filter.Age})
	}
	if filter.Gender != nil {
		query = query.Where(squirrel.Eq{"gender": *filter.Gender})
	}
	if filter.Breed != nil {
		query = query.Where(squirrel.Eq{"breed": *filter.Breed})
	}
	if filter.Size != nil {
		query = query.Where(squirrel.GtOrEq{"size": *filter.Size})
	}
	if filter.Weight != nil {
		query = query.Where(squirrel.GtOrEq{"weight": *filter.Weight})
	}

	query = query.OrderBy("id ASC").
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
		From("dogs").
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
		From("dogs").
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
	query := r.builder.Select("*").From("dogs")

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

func (r *PetRepository) AddPet(ctx context.Context, pet *entity.Pet) error {
	query := r.builder.Insert("dogs").
		Columns("name, gender, breed, wool_length, color, personality, size, age, weight, health, experience").
		Values(pet.Name, pet.Gender, pet.Breed, pet.WoolLength, pet.Color, pet.Personality, pet.Size, pet.Age, pet.Weight, pet.Health, pet.Experience)

	sql, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("get sql query: %w", err)
	}

	_, err = r.conn.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("save pet: %w", err)
	}

	return nil
}
