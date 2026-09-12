package repository

import (
	"context"

	db "github.com/henrystream/eduflex/school-service/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type SchoolRepository struct {
	queries *db.Queries
}

func NewSchoolRepository(q *db.Queries) *SchoolRepository {
	return &SchoolRepository{queries: q}
}

type CreateSchoolParams struct {
	Name         string
	Address      pgtype.Text
	ContactEmail pgtype.Text
	ContactPhone pgtype.Text
}

type UpdateSchoolParams struct {
	ID           pgtype.UUID
	Name         string
	Address      pgtype.Text
	ContactEmail pgtype.Text
	ContactPhone pgtype.Text
}

type DeleteSchoolParams struct {
	ID           pgtype.UUID
	Name         string
	Address      pgtype.Text
	ContactEmail pgtype.Text
	ContactPhone pgtype.Text
}

func (r *SchoolRepository) CreateSchool(ctx context.Context, p CreateSchoolParams) (db.School, error) {
	return r.queries.CreateSchool(ctx, db.CreateSchoolParams{
		Name:         p.Name,
		Address:      p.Address,
		ContactEmail: p.ContactEmail,
		ContactPhone: p.ContactPhone,
	})
}

func (r *SchoolRepository) GetSchool(ctx context.Context, id pgtype.UUID) (db.School, error) {
	return r.queries.GetSchool(ctx, id)
}

func (r *SchoolRepository) ListSchools(ctx context.Context) ([]db.School, error) {
	return r.queries.ListSchools(ctx)
}

func (r *SchoolRepository) UpdateSchool(ctx context.Context, p UpdateSchoolParams) (db.School, error) {
	return r.queries.UpdateSchool(ctx, db.UpdateSchoolParams{
		ID:           p.ID,
		Name:         p.Name,
		Address:      p.Address,
		ContactEmail: p.ContactEmail,
		ContactPhone: p.ContactPhone,
	})
}

func (r *SchoolRepository) DeleteSchool(ctx context.Context, p DeleteSchoolParams) error {
	return r.queries.DeleteSchool(ctx, p.ID)
}
