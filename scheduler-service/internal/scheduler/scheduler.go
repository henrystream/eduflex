package scheduler

import (
	"fmt"
	"time"

	"github.com/henrystream/eduflex/scheduler-service/internal/service"
)

type Scheduler struct {
	overdue   *service.OverdueService
	generator *service.InstallmentGenerationService
	interval  int
}

func NewScheduler(overdue *service.OverdueService, gen *service.InstallmentGenerationService, interval int) *Scheduler {
	return &Scheduler{
		overdue:   overdue,
		generator: gen,
		interval:  interval,
	}
}

func (s *Scheduler) Start(financingIDs []string) {
	for {
		fmt.Println("Scheduler running...")

		for _, id := range financingIDs {
			s.overdue.CheckOverdue(id)
			s.generator.GenerateMonthly(id)
		}

		time.Sleep(time.Duration(s.interval) * time.Second)
	}
}
