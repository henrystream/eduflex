package provider

import "fmt"

type SMSProvider struct {
	From string
}

func NewSMSProvider(from string) *SMSProvider {
	return &SMSProvider{From: from}
}

func (p *SMSProvider) SendSMS(to, message string) error {
	fmt.Printf("[SMS] From: %s | To: %s | Message: %s\n",
		p.From, to, message)
	return nil
}
