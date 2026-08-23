package service

import (
	"context"
	"errors"
	"math/big"

	db "github.com/henrystream/eduflex/financing-service/db/sqlc"
	"github.com/henrystream/eduflex/financing-service/internal/repository"
	"github.com/jackc/pgx/v5/pgtype"
)

type AgreementService struct {
	repo         *repository.AgreementRepository
	installments *InstallmentService
}

func NewAgreementService(r *repository.AgreementRepository, i *InstallmentService) *AgreementService {
	return &AgreementService{repo: r, installments: i}
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
		Status:       "ACTIVE",
	})
	if err != nil {
		return db.FinancingAgreement{}, err
	}

	// Generate installments
	err = s.installments.GenerateInstallments(ctx, agreement)
	if err != nil {
		return db.FinancingAgreement{}, err
	}

	return agreement, nil
}
