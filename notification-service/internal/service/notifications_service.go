package service

import (
	"fmt"

	"github.com/henrystream/eduflex/notification-service/internal/client"
	"github.com/henrystream/eduflex/notification-service/internal/provider"
)

type NotificationService struct {
	email *provider.EmailProvider
	sms   *provider.SMSProvider
}

func NewNotificationService(email *provider.EmailProvider, sms *provider.SMSProvider) *NotificationService {
	return &NotificationService{email: email, sms: sms}
}

func (s *NotificationService) ProcessEvent(ev client.DomainEvent) error {
	fmt.Println("NotificationService processing:", ev.EventType)

	switch ev.EventType {

	case "FINANCING_AGREEMENT_CREATED":
		s.notifyFinancingAgreement(ev)

	case "INSTALLMENT_CREATED":
		s.notifyInstallment(ev)

	case "LOAN_DRAWDOWN_CREATED":
		s.notifyDrawdown(ev)

	case "LOAN_REPAYMENT_CREATED":
		s.notifyRepayment(ev)

	case "STUDENT_PAYMENT_CREATED":
		s.notifyStudentPayment(ev)

	case "DISBURSEMENT_CREATED":
		s.notifyDisbursement(ev)

	default:
		fmt.Println("Unknown event type:", ev.EventType)
	}

	return nil
}

func (s *NotificationService) notifyFinancingAgreement(ev client.DomainEvent) {
	s.email.SendEmail("finance@eduflex.com", "New Financing Agreement", string(ev.Payload))
}

func (s *NotificationService) notifyInstallment(ev client.DomainEvent) {
	s.email.SendEmail("billing@eduflex.com", "New Installment Generated", string(ev.Payload))
}

func (s *NotificationService) notifyDrawdown(ev client.DomainEvent) {
	s.sms.SendSMS("+971500000000", "Loan drawdown executed")
}

func (s *NotificationService) notifyRepayment(ev client.DomainEvent) {
	s.sms.SendSMS("+971500000000", "Loan repayment received")
}

func (s *NotificationService) notifyStudentPayment(ev client.DomainEvent) {
	s.email.SendEmail("student@eduflex.com", "Payment Received", string(ev.Payload))
}

func (s *NotificationService) notifyDisbursement(ev client.DomainEvent) {
	s.email.SendEmail("school@eduflex.com", "Disbursement Completed", string(ev.Payload))
}
