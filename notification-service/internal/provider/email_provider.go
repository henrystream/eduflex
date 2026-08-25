package provider

import "fmt"

type EmailProvider struct {
	From string
}

func NewEmailProvider(from string) *EmailProvider {
	return &EmailProvider{From: from}
}

func (p *EmailProvider) SendEmail(to, subject, body string) error {
	fmt.Printf("[EMAIL] From: %s | To: %s | Subject: %s | Body: %s\n",
		p.From, to, subject, body)
	return nil
}
