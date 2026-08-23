package service

import (
	"fmt"

	"github.com/henrystream/eduflex/reporting-service/internal/client"
	"github.com/shopspring/decimal"
)

type FinancialReportService struct {
	ledger *client.LedgerClient
}

func NewFinancialReportService(ledger *client.LedgerClient) *FinancialReportService {
	return &FinancialReportService{ledger: ledger}
}

type FinancialStatement struct {
	TotalRevenue       string `json:"total_revenue"`
	TotalDrawdowns     string `json:"total_drawdowns"`
	TotalRepayments    string `json:"total_repayments"`
	TotalDisbursements string `json:"total_disbursements"`
	Cashflow           string `json:"cashflow"`
	Receivables        string `json:"receivables"`
	Payables           string `json:"payables"`
}

func (s *FinancialReportService) GenerateFinancialStatement() (FinancialStatement, error) {
	entries, err := s.ledger.ListLedgerEntries()
	if err != nil {
		return FinancialStatement{}, fmt.Errorf("failed to list ledger entries: %w", err)
	}

	revenue := SumLedger(entries, "INSTALLMENT")
	drawdowns := SumLedger(entries, "LOAN_DRAWDOWN")
	repayments := SumLedger(entries, "LOAN_REPAYMENT")
	disbursements := SumLedger(entries, "DISBURSEMENT")

	cashflow := Subtract(repayments, disbursements)
	receivables := SumLedger(entries, "FINANCING_AGREEMENT")
	payables := disbursements

	return FinancialStatement{
		TotalRevenue:       revenue,
		TotalDrawdowns:     drawdowns,
		TotalRepayments:    repayments,
		TotalDisbursements: disbursements,
		Cashflow:           cashflow,
		Receivables:        receivables,
		Payables:           payables,
	}, nil
}

// SumLedger sums all amounts in ledger entries filtered by event type
func SumLedger(entries []client.LedgerEntry, eventType string) string {
	if len(entries) == 0 {
		return "0"
	}

	total := decimal.Zero
	for _, entry := range entries {
		if entry.EventType == eventType && entry.Amount.Valid {
			val := pgTypeNumericToDecimal(entry.Amount)
			total = total.Add(val)
		}
	}

	return total.String()
}
