package service

import (
	"fmt"

	"github.com/henrystream/eduflex/reporting-service/internal/client"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shopspring/decimal"
)

type SchoolReportService struct {
	ledger  *client.LedgerClient
	student *client.StudentClient
	school  *client.SchoolClient
}

func NewSchoolReportService(ledger *client.LedgerClient, student *client.StudentClient, school *client.SchoolClient) *SchoolReportService {
	return &SchoolReportService{ledger: ledger, student: student, school: school}
}

type SchoolStatement struct {
	SchoolInfo       any    `json:"school_info"`
	TotalDisbursed   string `json:"total_disbursed"`
	TotalStudentPaid string `json:"total_student_paid"`
	Outstanding      string `json:"outstanding"`
	LedgerEntries    any    `json:"ledger_entries"`
}

func (s *SchoolReportService) GenerateSchoolStatement(schoolID string) (SchoolStatement, error) {
	schoolInfo, err := s.school.GetSchool(schoolID)
	if err != nil {
		return SchoolStatement{}, fmt.Errorf("failed to get school info: %w", err)
	}

	disbursements, err := s.student.ListDisbursementsBySchool(schoolID)
	if err != nil {
		return SchoolStatement{}, fmt.Errorf("failed to list disbursements: %w", err)
	}

	payments, err := s.student.ListPaymentsBySchool(schoolID)
	if err != nil {
		return SchoolStatement{}, fmt.Errorf("failed to list payments: %w", err)
	}

	ledgerEntries, err := s.ledger.ListByService("disbursement-service")
	if err != nil {
		return SchoolStatement{}, fmt.Errorf("failed to list ledger entries: %w", err)
	}

	totalDisbursed := SumAmounts(disbursements)
	totalPaid := SumAmounts(payments)
	outstanding := Subtract(totalDisbursed, totalPaid)

	return SchoolStatement{
		SchoolInfo:       schoolInfo,
		TotalDisbursed:   totalDisbursed,
		TotalStudentPaid: totalPaid,
		Outstanding:      outstanding,
		LedgerEntries:    ledgerEntries,
	}, nil
}

// SumAmounts adds up all amounts in the given slice
func SumAmounts(amounts []pgtype.Numeric) string {
	if len(amounts) == 0 {
		return "0"
	}

	total := decimal.Zero
	for _, amount := range amounts {
		if amount.Valid && !amount.NaN {
			val := pgTypeNumericToDecimal(amount)
			total = total.Add(val)
		}
	}

	return total.String()
}

// pgTypeNumericToDecimal converts pgtype.Numeric to decimal.Decimal
func pgTypeNumericToDecimal(n pgtype.Numeric) decimal.Decimal {
	if !n.Valid || n.NaN || n.Int == nil {
		return decimal.Zero
	}

	d := decimal.NewFromBigInt(n.Int, -n.Exp)
	return d
}

// Subtract calculates the difference between two numeric amounts
func Subtract(a, b string) string {
	decimalA, err := decimal.NewFromString(a)
	if err != nil {
		return "0"
	}

	decimalB, err := decimal.NewFromString(b)
	if err != nil {
		return "0"
	}

	result := decimalA.Sub(decimalB)
	return result.String()
}
