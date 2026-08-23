package service

import (
	"context"
	"errors"

	db "github.com/henrystream/eduflex/loan-service/db/sqlc"
	"github.com/henrystream/eduflex/loan-service/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type RepaymentService struct {
	repo *repository.RepaymentRepository
}

func NewRepaymentService(r *repository.RepaymentRepository) *RepaymentService {
	return &RepaymentService{repo: r}
}

type CreateRepaymentRequest struct {
	DrawdownID pgtype.UUID    `json:"drawdown_id"`
	Amount     pgtype.Numeric `json:"amount"`
	Reference  string         `json:"reference"`
}

func (s *RepaymentService) CreateRepayment(ctx context.Context, req CreateRepaymentRequest) (db.LoanRepayment, error) {
	if req.DrawdownID.String() == "" {
		return db.LoanRepayment{}, errors.New("drawdown_id required")
	}

	return s.repo.CreateRepayment(ctx, repository.CreateRepaymentParams{
		DrawdownID: req.DrawdownID,
		Amount:     req.Amount,
		Reference:  req.Reference,
	})
}

func (s *RepaymentService) ListRepaymentsByDrawdown(ctx context.Context, drawdownID pgtype.UUID) ([]db.LoanRepayment, error) {
	return s.repo.ListRepaymentsByDrawdown(ctx, drawdownID)
}
