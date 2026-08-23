package service

import (
	"context"
	"errors"
	"time"

	db "github.com/henrystream/eduflex/disbursement-service/db/sqlc"
	"github.com/henrystream/eduflex/disbursement-service/internal/ledger"
	"github.com/henrystream/eduflex/disbursement-service/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type DisbursementService struct {
	repo   *repository.DisbursementRepository
	ledger *ledger.LedgerClient
}

func NewDisbursementService(r *repository.DisbursementRepository, l *ledger.LedgerClient) *DisbursementService {
	return &DisbursementService{repo: r, ledger: l}
}

type CreateDisbursementRequest struct {
	SchoolID      pgtype.UUID    `json:"school_id"`
	InvoiceID     pgtype.UUID    `json:"invoice_id"`
	Amount        pgtype.Numeric `json:"amount"`
	PaymentMethod string         `json:"payment_method"`
	Reference     string         `json:"reference"`
}

func (s *DisbursementService) CreateDisbursement(ctx context.Context, req CreateDisbursementRequest) (db.EduflexDisbursement, error) {
	if req.SchoolID.String() == "" || req.InvoiceID.String() == "" {
		return db.EduflexDisbursement{}, errors.New("school_id and invoice_id required")
	}

	disbursement, err := s.repo.CreateDisbursement(ctx, repository.CreateDisbursementParams{
		SchoolID:      req.SchoolID,
		InvoiceID:     req.InvoiceID,
		Amount:        req.Amount,
		PaymentMethod: req.PaymentMethod,
		Reference:     req.Reference,
		Status:        "COMPLETED",
	})
	if err != nil {
		return db.EduflexDisbursement{}, err
	}

	if s.ledger != nil {
		if err := s.ledger.CreateEntry(ledger.LedgerEntryRequest{
			EventType:     "DISBURSEMENT",
			EventID:       disbursement.ID,
			SourceService: "disbursement-service",
			DebitAccount:  "School Payables",
			CreditAccount: "Cash - Bank",
			Amount:        disbursement.Amount,
			Currency:      "AED",
			OccurredAt: pgtype.Timestamptz{
				Time:  time.Now(),
				Valid: true,
			},
		}); err != nil {
			return db.EduflexDisbursement{}, err
		}
	}

	return disbursement, nil
}

func (s *DisbursementService) ListDisbursementsBySchool(ctx context.Context, schoolID pgtype.UUID) ([]db.EduflexDisbursement, error) {
	return s.repo.ListDisbursementsBySchool(ctx, schoolID)
}
