package service

import (
	"context"
	"errors"

	db "github.com/henrystream/eduflex/loan-service/db/sqlc"
	"github.com/henrystream/eduflex/loan-service/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type BankService struct {
	repo *repository.BankRepository
}

func NewBankService(r *repository.BankRepository) *BankService {
	return &BankService{repo: r}
}

type CreateBankRequest struct {
	Name         string      `json:"name"`
	ContactEmail pgtype.Text `json:"contact_email"`
	ContactPhone pgtype.Text `json:"contact_phone"`
}

func (s *BankService) CreateBank(ctx context.Context, req CreateBankRequest) (db.Bank, error) {
	if req.Name == "" {
		return db.Bank{}, errors.New("name is required")
	}
	return s.repo.CreateBank(ctx, req.Name, req.ContactEmail, req.ContactPhone)
}

func (s *BankService) ListBanks(ctx context.Context) ([]db.Bank, error) {
	return s.repo.ListBanks(ctx)
}
