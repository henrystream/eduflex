package service

import (
	"github.com/henrystream/eduflex/reconciliation-service/internal/client"
	"github.com/henrystream/eduflex/reconciliation-service/internal/parser"
)

type MatchResult struct {
	BankTx   parser.BankTransaction
	LedgerTx *client.LedgerEntry
	Status   string // MATCHED, MISMATCHED, MISSING
}

type MatchingEngine struct{}

func NewMatchingEngine() *MatchingEngine {
	return &MatchingEngine{}
}

func (m *MatchingEngine) Match(bankTxs []parser.BankTransaction, ledgerTxs []client.LedgerEntry) []MatchResult {
	results := []MatchResult{}

	ledgerMap := map[string]client.LedgerEntry{}
	for _, l := range ledgerTxs {
		ledgerMap[l.CreditAccount+"-"+l.Amount] = l
		ledgerMap[l.DebitAccount+"-"+l.Amount] = l
	}

	for _, b := range bankTxs {
		key := b.Type + "-" + b.Amount

		if l, ok := ledgerMap[key]; ok {
			results = append(results, MatchResult{
				BankTx:   b,
				LedgerTx: &l,
				Status:   "MATCHED",
			})
		} else {
			results = append(results, MatchResult{
				BankTx:   b,
				LedgerTx: nil,
				Status:   "MISSING",
			})
		}
	}

	return results
}
