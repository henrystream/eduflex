package service

import (
	"fmt"

	"github.com/henrystream/eduflex/worker-service/internal/client"
)

type ProcessorService struct{}

func NewProcessorService() *ProcessorService {
	return &ProcessorService{}
}

func (p *ProcessorService) Process(ev client.DomainEvent) error {
	fmt.Println("Processing event:", ev.EventType)

	switch ev.EventType {

	case "FINANCING_AGREEMENT_CREATED":
		p.handleFinancingAgreement(ev)

	case "INSTALLMENT_CREATED":
		p.handleInstallment(ev)

	case "LOAN_DRAWDOWN_CREATED":
		p.handleDrawdown(ev)

	case "LOAN_REPAYMENT_CREATED":
		p.handleRepayment(ev)

	case "STUDENT_PAYMENT_CREATED":
		p.handleStudentPayment(ev)

	case "DISBURSEMENT_CREATED":
		p.handleDisbursement(ev)

	default:
		fmt.Println("Unknown event type:", ev.EventType)
	}

	return nil
}

func (p *ProcessorService) handleFinancingAgreement(ev client.DomainEvent) {
	fmt.Println("Triggering financing agreement notifications")
}

func (p *ProcessorService) handleInstallment(ev client.DomainEvent) {
	fmt.Println("Triggering installment recalculation")
}

func (p *ProcessorService) handleDrawdown(ev client.DomainEvent) {
	fmt.Println("Triggering disbursement workflow")
}

func (p *ProcessorService) handleRepayment(ev client.DomainEvent) {
	fmt.Println("Triggering repayment ledger sync")
}

func (p *ProcessorService) handleStudentPayment(ev client.DomainEvent) {
	fmt.Println("Triggering student payment confirmation")
}

func (p *ProcessorService) handleDisbursement(ev client.DomainEvent) {
	fmt.Println("Triggering school payout notification")
}
