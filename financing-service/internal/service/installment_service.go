package service

import (
	"context"

	"math/big"
	"time"

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
	total, _ := new(big.Float).SetString(agreement.TotalPayable.InfinityModifier.String())
	months := agreement.TermMonths

	monthly := new(big.Float).Quo(total, big.NewFloat(float64(months)))
	monthlyStr := monthly.Text('f', 2)
	var ms pgtype.Numeric
	ms.Scan(monthlyStr)

	start, _ := time.Parse("2006-01-02", agreement.StartDate.Time.String())

	for i := int32(1); i <= months; i++ {
		due := start.AddDate(0, int(i), 0)
		var duedate pgtype.Date
		duedate.Scan(due)
		_, err := s.repo.CreateInstallment(ctx, repository.CreateInstallmentParams{
			FinancingID:       agreement.ID,
			InstallmentNumber: i,
			DueDate:           duedate, //due.Format("2006-01-02"),
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
