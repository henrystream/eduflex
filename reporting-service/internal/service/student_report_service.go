package service

import (
	"fmt"

	"github.com/henrystream/eduflex/reporting-service/internal/client"
)

type StudentReportService struct {
	ledger    *client.LedgerClient
	financing *client.FinancingClient
	student   *client.StudentClient
}

func NewStudentReportService(ledger *client.LedgerClient, financing *client.FinancingClient, student *client.StudentClient) *StudentReportService {
	return &StudentReportService{ledger: ledger, financing: financing, student: student}
}

type StudentLoanSummary struct {
	StudentInfo   any    `json:"student_info"`
	Agreements    any    `json:"agreements"`
	Installments  any    `json:"installments"`
	Payments      any    `json:"payments"`
	Outstanding   string `json:"outstanding"`
	LedgerEntries any    `json:"ledger_entries"`
}

func (s *StudentReportService) GenerateStudentLoanSummary(studentID string) (StudentLoanSummary, error) {
	studentInfo, err := s.student.GetStudent(studentID)
	if err != nil {
		return StudentLoanSummary{}, fmt.Errorf("failed to get student info: %w", err)
	}

	agreements, err := s.financing.ListAgreementsByStudent(studentID)
	if err != nil {
		return StudentLoanSummary{}, fmt.Errorf("failed to list agreements: %w", err)
	}

	details, installments, err := s.financing.ListInstallmentsByStudent(studentID)
	if err != nil {
		return StudentLoanSummary{}, fmt.Errorf("failed to list installments: %w", err)
	}

	payments, err := s.student.ListPaymentsByStudent(studentID)
	if err != nil {
		return StudentLoanSummary{}, fmt.Errorf("failed to list payments: %w", err)
	}

	ledgerEntries, err := s.ledger.ListByService("student-service")
	if err != nil {
		return StudentLoanSummary{}, fmt.Errorf("failed to list ledger entries: %w", err)
	}

	totalInstallments := SumAmounts(installments)
	totalPaid := SumAmounts(payments)
	outstanding := Subtract(totalInstallments, totalPaid)

	return StudentLoanSummary{
		StudentInfo:   studentInfo,
		Agreements:    agreements,
		Installments:  details, //installments,
		Payments:      payments,
		Outstanding:   outstanding,
		LedgerEntries: ledgerEntries,
	}, nil
}
