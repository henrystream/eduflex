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

	principal, _ := new(big.Float).SetString(req.Principal.InfinityModifier.String())
	interestRate, _ := new(big.Float).SetString(req.InterestRate.InfinityModifier.String())
	serviceFee, _ := new(big.Float).SetString(req.ServiceFee.InfinityModifier.String())

	// Calculate interest
	interest := new(big.Float).Mul(principal, interestRate)

	// Total payable = principal + interest + service fee
	totalPayable := new(big.Float).Add(principal, interest)
	totalPayable.Add(totalPayable, serviceFee)

	totalPayableStr := totalPayable.Text('f', 2)

	var tp pgtype.Numeric
	tp.Scan(totalPayableStr)

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
