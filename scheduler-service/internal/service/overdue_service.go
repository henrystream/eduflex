package service

import (
	"fmt"
	"time"

	"github.com/henrystream/eduflex/scheduler-service/internal/client"
)

type OverdueService struct {
	financing *client.FinancingClient
	events    *client.EventsClient
}

func NewOverdueService(fin *client.FinancingClient, ev *client.EventsClient) *OverdueService {
	return &OverdueService{financing: fin, events: ev}
}

func (s *OverdueService) CheckOverdue(financingID string) error {
	installments, err := s.financing.ListInstallments(financingID)
	if err != nil {
		return err
	}

	today := time.Now()

	for _, inst := range installments {
		due, _ := time.Parse("2006-01-02", inst.DueDate)

		if inst.Status == "PENDING" && today.After(due) {
			fmt.Println("Installment overdue:", inst.ID)

			s.events.Publish(
				"INSTALLMENT_OVERDUE",
				inst.ID,
				today.Format("2006-01-02"),
				inst,
			)
		}
	}

	return nil
}
