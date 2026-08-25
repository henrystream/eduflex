package worker

import (
	"fmt"
	"time"

	"github.com/henrystream/eduflex/notification-service/internal/client"
	"github.com/henrystream/eduflex/notification-service/internal/service"
)

type NotificationWorker struct {
	eventsClient *client.EventsClient
	notifier     *service.NotificationService
	interval     int
}

func NewNotificationWorker(ec *client.EventsClient, n *service.NotificationService, interval int) *NotificationWorker {
	return &NotificationWorker{
		eventsClient: ec,
		notifier:     n,
		interval:     interval,
	}
}

func (w *NotificationWorker) Start() {
	for {
		fmt.Println("Notification worker polling...")

		events, err := w.eventsClient.ListUnprocessed()
		if err != nil {
			fmt.Println("Error fetching events:", err)
			time.Sleep(time.Duration(w.interval) * time.Second)
			continue
		}

		for _, ev := range events {
			fmt.Println("Sending notifications for:", ev.EventType)

			if err := w.notifier.ProcessEvent(ev); err != nil {
				fmt.Println("Error processing event:", err)
				continue
			}

			w.eventsClient.MarkProcessed(ev.ID)
		}

		time.Sleep(time.Duration(w.interval) * time.Second)
	}
}
