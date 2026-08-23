package service

import (
	"context"
	"errors"

	"math/big"

	db "github.com/henrystream/eduflex/financing-service/db/sqlc"
	"github.com/henrystream/eduflex/financing-service/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type InstallmentService struct {
	repo *repository.InstallmentRepository
}

func NewInstallmentService(r *repository.InstallmentRepository) *InstallmentService {
	return &InstallmentService{repo: r}
}

func (s *InstallmentService) GenerateInstallments(ctx context.Context, agreement db.FinancingAgreement) error {
	total, err := numericToBigFloat(agreement.TotalPayable)
	if err != nil {
		return err
	}
	months := agreement.TermMonths
	if months <= 0 {
		return errors.New("term_months must be greater than zero")
	}

	monthly := new(big.Float).Quo(total, big.NewFloat(float64(months)))
	monthlyStr := monthly.Text('f', 2)
	var ms pgtype.Numeric
	if err := ms.Scan(monthlyStr); err != nil {
		return err
	}

	if !agreement.StartDate.Valid {
		return errors.New("start_date is required")
	}
	start := agreement.StartDate.Time

	for i := int32(1); i <= months; i++ {
		due := start.AddDate(0, int(i), 0)
		var dueDate pgtype.Date
		if err := dueDate.Scan(due); err != nil {
			return err
		}
		_, err := s.repo.CreateInstallment(ctx, repository.CreateInstallmentParams{
			FinancingID:       agreement.ID,
			InstallmentNumber: i,
			DueDate:           dueDate,
			Amount:            ms,
			Status:            "PENDING",
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *InstallmentService) ListInstallments(ctx context.Context, financingID pgtype.UUID) ([]db.MonthlyInstallment, error) {
	return s.repo.ListInstallments(ctx, financingID)
}
