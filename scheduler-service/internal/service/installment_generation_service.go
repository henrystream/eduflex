package service

import (
	"fmt"
	"time"

	"github.com/henrystream/eduflex/scheduler-service/internal/client"
)

type InstallmentGenerationService struct {
	financing *client.FinancingClient
	events    *client.EventsClient
}

func NewInstallmentGenerationService(fin *client.FinancingClient, ev *client.EventsClient) *InstallmentGenerationService {
	return &InstallmentGenerationService{financing: fin, events: ev}
}

func (s *InstallmentGenerationService) GenerateMonthly(financingID string) error {
	installments, err := s.financing.ListInstallments(financingID)
	if err != nil {
		return err
	}

	today := time.Now()

	for _, inst := range installments {
		due, _ := time.Parse("2006-01-02", inst.DueDate)

		if due.Year() == today.Year() && due.Month() == today.Month() && inst.Status == "PENDING" {
			fmt.Println("Installment due today:", inst.ID)

			s.events.Publish(
				"INSTALLMENT_DUE_TODAY",
				inst.ID,
				today.Format("2006-01-02"),
				inst,
			)
		}
	}

	return nil
}
