package repository

import (
	"context"

	db "github.com/henrystream/eduflex/loan-service/db/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type BankRepository struct {
	queries *db.Queries
}

func NewBankRepository(q *db.Queries) *BankRepository {
	return &BankRepository{queries: q}
}

func (r *BankRepository) CreateBank(ctx context.Context, name string, email, phone pgtype.Text) (db.Bank, error) {
	return r.queries.CreateBank(ctx, db.CreateBankParams{
		Name:         name,
		ContactEmail: email,
		ContactPhone: phone,
	})
}

func (r *BankRepository) ListBanks(ctx context.Context) ([]db.Bank, error) {
	return r.queries.ListBanks(ctx)
}
