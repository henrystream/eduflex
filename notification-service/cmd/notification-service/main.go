package main

import (
	"github.com/henrystream/eduflex/notification-service/internal/client"
	"github.com/henrystream/eduflex/notification-service/internal/config"
	"github.com/henrystream/eduflex/notification-service/internal/provider"
	"github.com/henrystream/eduflex/notification-service/internal/service"
	"github.com/henrystream/eduflex/notification-service/internal/worker"
)

func main() {
	cfg := config.Load()

	eventsClient := client.NewEventsClient(cfg.EventsURL)
	emailProvider := provider.NewEmailProvider(cfg.EmailFrom)
	smsProvider := provider.NewSMSProvider(cfg.SMSFrom)

	notifier := service.NewNotificationService(emailProvider, smsProvider)
	w := worker.NewNotificationWorker(eventsClient, notifier, cfg.Interval)

	w.Start()
}
