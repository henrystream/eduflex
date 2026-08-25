package worker

import (
	"fmt"
	"time"

	"github.com/henrystream/eduflex/worker-service/internal/client"
	"github.com/henrystream/eduflex/worker-service/internal/service"
)

type EventWorker struct {
	eventsClient *client.EventsClient
	processor    *service.ProcessorService
	interval     int
}

func NewEventWorker(ec *client.EventsClient, ps *service.ProcessorService, interval int) *EventWorker {
	return &EventWorker{
		eventsClient: ec,
		processor:    ps,
		interval:     interval,
	}
}

func (w *EventWorker) Start() {
	for {
		fmt.Println("Worker polling for events...")

		events, err := w.eventsClient.ListUnprocessed()
		if err != nil {
			fmt.Println("Error fetching events:", err)
			time.Sleep(time.Duration(w.interval) * time.Second)
			continue
		}

		for _, ev := range events {
			fmt.Println("Processing event:", ev.ID)

			if err := w.processor.Process(ev); err != nil {
				fmt.Println("Error processing event:", err)
				continue
			}

			w.eventsClient.MarkProcessed(ev.ID)
		}

		time.Sleep(time.Duration(w.interval) * time.Second)
	}
}
