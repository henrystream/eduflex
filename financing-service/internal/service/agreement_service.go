package service

import (
	"context"
	"errors"
	"math/big"
	"time"

	db "github.com/henrystream/eduflex/financing-service/db/sqlc"
	"github.com/henrystream/eduflex/financing-service/internal/events"
	"github.com/henrystream/eduflex/financing-service/internal/fraud"
	"github.com/henrystream/eduflex/financing-service/internal/repository"

	"github.com/jackc/pgx/v5/pgtype"
)

type AgreementService struct {
	repo         *repository.AgreementRepository
	installments *InstallmentService
	events       *events.EventClient
	fraud        *fraud.FraudClient
	ledger       LedgerClient
}

type LedgerClient interface {
	CreateEntry(req LedgerEntryRequest) error
}

type LedgerEntryRequest struct {
	EventType     string             `json:"event_type"`
	EventID       pgtype.UUID        `json:"event_id"`
	SourceService string             `json:"source_service"`
	DebitAccount  string             `json:"debit_account"`
	CreditAccount string             `json:"credit_account"`
	Amount        pgtype.Numeric     `json:"amount"`
	Currency      string             `json:"currency"`
	OccurredAt    pgtype.Timestamptz `json:"occurred_at"`
}

func NewAgreementService(
	repo *repository.AgreementRepository,
	installments *InstallmentService,
	events *events.EventClient,
	fraud *fraud.FraudClient,
	ledger ...LedgerClient,
) *AgreementService {
	var ledgerClient LedgerClient
	if len(ledger) > 0 {
		ledgerClient = ledger[0]
	}
	return &AgreementService{
		repo:         repo,
		installments: installments,
		events:       events,
		fraud:        fraud,
		ledger:       ledgerClient,
	}
}

type CreateAgreementRequest struct {
	StudentID    pgtype.UUID    `json:"student_id"`
	InvoiceID    pgtype.UUID    `json:"invoice_id"`
	Principal    pgtype.Numeric `json:"principal"`
	InterestRate pgtype.Numeric `json:"interest_rate"`
	ServiceFee   pgtype.Numeric `json:"service_fee"`
	TermMonths   int32          `json:"term_months"`
	StartDate    pgtype.Date    `json:"start_date"`
}

func (s *AgreementService) CreateAgreement(ctx context.Context, req CreateAgreementRequest) (db.FinancingAgreement, error) {
	if req.StudentID.String() == "" || req.InvoiceID.String() == "" {
		return db.FinancingAgreement{}, errors.New("student_id and invoice_id required")
	}
	if req.TermMonths <= 0 {
		return db.FinancingAgreement{}, errors.New("term_months must be greater than zero")
	}

	principal, err := numericToBigFloat(req.Principal)
	if err != nil {
		return db.FinancingAgreement{}, errors.New("invalid principal")
	}
	interestRate, err := numericToBigFloat(req.InterestRate)
	if err != nil {
		return db.FinancingAgreement{}, errors.New("invalid interest_rate")
	}
	serviceFee, err := numericToBigFloat(req.ServiceFee)
	if err != nil {
		return db.FinancingAgreement{}, errors.New("invalid service_fee")
	}

	// Calculate interest
	interest := new(big.Float).Mul(principal, interestRate)

	// Total payable = principal + interest + service fee
	totalPayable := new(big.Float).Add(principal, interest)
	totalPayable.Add(totalPayable, serviceFee)

	totalPayableStr := totalPayable.Text('f', 2)

	var tp pgtype.Numeric
	if err := tp.Scan(totalPayableStr); err != nil {
		return db.FinancingAgreement{}, err
	}

	agreement, err := s.repo.CreateAgreement(ctx, repository.CreateAgreementParams{
		StudentID:    req.StudentID,
		InvoiceID:    req.InvoiceID,
		Principal:    req.Principal,
		InterestRate: req.InterestRate,
		ServiceFee:   req.ServiceFee,
		TotalPayable: tp,
		TermMonths:   req.TermMonths,
		StartDate:    req.StartDate,
		Status:       "PENDING",
	})
	if err != nil {
		return db.FinancingAgreement{}, err
	}
	// 3. FRAUD CHECK
	fraudResult, _ := s.fraud.CheckAgreement(agreement.ID)

	// 4. BLOCK HIGH-RISK AGREEMENTS
	if fraudResult.AgreementRisk.Level == "HIGH" ||
		fraudResult.StudentRisk.Level == "HIGH" ||
		fraudResult.SchoolRisk.Level == "HIGH" {

		// Mark agreement as REJECTED
		s.repo.UpdateStatus(ctx, agreement.ID, "REJECTED_FRAUD")

		// Publish event
		s.events.Publish(
			"AGREEMENT_REJECTED_FRAUD",
			"financing-service",
			agreement.ID,
			pgtype.Timestamptz{Time: time.Now(), Valid: true},
			fraudResult,
		)

		return db.FinancingAgreement{}, errors.New("agreement rejected due to high fraud risk")
	}

	// 5. APPROVE AGREEMENT
	approved, err := s.repo.UpdateStatus(ctx, agreement.ID, "APPROVED")
	if err != nil {
		return db.FinancingAgreement{}, err
	}

	// 6. Publish approval event
	s.events.Publish(
		"FINANCING_AGREEMENT_APPROVED",
		"financing-service",
		agreement.ID,
		pgtype.Timestamptz{Time: time.Now(), Valid: true},
		approved,
	)

	// Generate installments
	err = s.installments.GenerateInstallments(ctx, approved)
	if err != nil {
		return db.FinancingAgreement{}, err
	}

	if s.ledger != nil {
		_ = s.ledger.CreateEntry(LedgerEntryRequest{
			EventType:     "FINANCING_AGREEMENT",
			EventID:       approved.ID,
			SourceService: "financing-service",
			DebitAccount:  "Accounts Receivable - Student",
			CreditAccount: "Financing Principal",
			Amount:        approved.Principal,
			Currency:      "AED",
			OccurredAt: pgtype.Timestamptz{
				Time:  approved.StartDate.Time,
				Valid: approved.StartDate.Valid,
			},
		})
	}

	var eventClient events.EventClient

	pubRequest := events.PublishEventRequest{
		EventType:     "FINANCING_AGREEMENT_CREATED",
		SourceService: "financing-service",
		AggregateID:   approved.ID,
		OccurredAt:    pgtype.Timestamptz{Time: approved.StartDate.Time, Valid: true},
		Payload:       approved,
	}

	err = eventClient.Publish(pubRequest.EventType, pubRequest.SourceService, pubRequest.AggregateID, pubRequest.OccurredAt, pubRequest.Payload)
	if err != nil {
		return approved, err
	}

	return approved, nil
}

func (s *AgreementService) ListAgreementsByStudent(ctx context.Context, studentID pgtype.UUID) ([]db.FinancingAgreement, error) {
	return s.repo.ListAgreementsByStudent(ctx, studentID)
}
