package service

import (
	"fmt"
	"time"

	"github.com/henrystream/eduflex/reconciliation-service/internal/client"
	"github.com/henrystream/eduflex/reconciliation-service/internal/parser"
)

type ReconciliationService struct {
	ledger  *client.LedgerClient
	events  *client.EventsClient
	matcher *MatchingEngine
}

func NewReconciliationService(l *client.LedgerClient, e *client.EventsClient, m *MatchingEngine) *ReconciliationService {
	return &ReconciliationService{ledger: l, events: e, matcher: m}
}

func (s *ReconciliationService) Reconcile(path string) ([]MatchResult, error) {
	bankTxs, err := parser.ParseCSV(path)
	if err != nil {
		return nil, err
	}

	ledgerTxs, err := s.ledger.ListAll()
	if err != nil {
		return nil, err
	}

	results := s.matcher.Match(bankTxs, ledgerTxs)

	for _, r := range results {
		if r.Status != "MATCHED" {
			fmt.Println("Publishing mismatch event:", r.BankTx.Reference)

			s.events.Publish(
				"RECONCILIATION_MISMATCH",
				r.BankTx.Reference,
				time.Now().Format("2006-01-02"),
				r,
			)
		}
	}

	return results, nil
}
