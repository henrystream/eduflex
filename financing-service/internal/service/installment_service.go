package service

import (
	"context"
	"errors"

	"math/big"

	db "github.com/henrystream/eduflex/financing-service/db/sqlc"
	"github.com/henrystream/eduflex/financing-service/internal/events"
	"github.com/henrystream/eduflex/financing-service/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type InstallmentService struct {
	repo   *repository.InstallmentRepository
	ledger LedgerClient
}

func NewInstallmentService(r *repository.InstallmentRepository, ledger ...LedgerClient) *InstallmentService {
	var ledgerClient LedgerClient
	if len(ledger) > 0 {
		ledgerClient = ledger[0]
	}
	return &InstallmentService{repo: r, ledger: ledgerClient}
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
		installment, err := s.repo.CreateInstallment(ctx, repository.CreateInstallmentParams{
			FinancingID:       agreement.ID,
			InstallmentNumber: i,
			DueDate:           dueDate,
			Amount:            ms,
			Status:            "PENDING",
		})
		if err != nil {
			return err
		}

		if s.ledger != nil {
			_ = s.ledger.CreateEntry(LedgerEntryRequest{
				EventType:     "INSTALLMENT",
				EventID:       installment.ID,
				SourceService: "financing-service",
				DebitAccount:  "Accounts Receivable - Installments",
				CreditAccount: "Installment Revenue",
				Amount:        installment.Amount,
				Currency:      "AED",
				OccurredAt: pgtype.Timestamptz{
					Time:  installment.DueDate.Time,
					Valid: installment.DueDate.Valid,
				},
			})
		}

		var eventClient events.EventClient

		pubRequest := events.PublishEventRequest{
			EventType:     "INSTALLMENT_CREATED",
			SourceService: "financing-service",
			AggregateID:   installment.ID,
			OccurredAt:    pgtype.Timestamptz{Time: installment.DueDate.Time, Valid: true},
			Payload:       installment,
		}

		eventClient.Publish(pubRequest.EventType, pubRequest.AggregateID, pubRequest.OccurredAt, pubRequest.Payload)

	}

	return nil
}

func (s *InstallmentService) ListInstallments(ctx context.Context, financingID pgtype.UUID) ([]db.MonthlyInstallment, error) {
	return s.repo.ListInstallments(ctx, financingID)
}
