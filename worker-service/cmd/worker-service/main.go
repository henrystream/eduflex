package main

import (
	"github.com/henrystream/eduflex/worker-service/internal/client"
	"github.com/henrystream/eduflex/worker-service/internal/config"
	"github.com/henrystream/eduflex/worker-service/internal/service"
	"github.com/henrystream/eduflex/worker-service/internal/worker"
)

func main() {
	cfg := config.Load()

	eventsClient := client.NewEventsClient(cfg.EventsURL)
	processor := service.NewProcessorService()
	w := worker.NewEventWorker(eventsClient, processor, cfg.Interval)

	w.Start()
}
