package main

import (
	"context"
	"log"
	"time"

	concurrencypatterns "github.com/arunsworld/concurrency-patterns/concurrency-patterns"
	"github.com/arunsworld/nursery"
)

func main() {
	start := time.Now()
	err := nursery.RunConcurrently(
		// 50 jobs with 100ms
		func(ctx context.Context, errCh chan error) {
			worker := concurrencypatterns.NewWorker()
			err := nursery.RunMultipleCopiesConcurrentlyWithContext(ctx, 50,
				func(ctx context.Context, errCh chan error) {
					if err := worker.Work(ctx, false); err != nil {
						errCh <- err
					}
				},
			)
			if err != nil {
				errCh <- err
			}
		},
		// 50 jobs with 1s
		func(ctx context.Context, errCh chan error) {
			worker := concurrencypatterns.NewWorker(concurrencypatterns.SetAveDelayOnWorker(1000))
			err := nursery.RunMultipleCopiesConcurrentlyWithContext(ctx, 50,
				func(ctx context.Context, errCh chan error) {
					if err := worker.Work(ctx, false); err != nil {
						errCh <- err
					}
				},
			)
			if err != nil {
				errCh <- err
			}
		},
		// 100ms job with error
		func(ctx context.Context, errCh chan error) {
			worker := concurrencypatterns.NewWorker()
			if err := worker.Work(ctx, true); err != nil {
				errCh <- err
			}
		},
	)
	if err == nil {
		log.Fatalf("no error was returned in %v", time.Since(start))
	}
	log.Printf("finished with error (%v) in: %v", err, time.Since(start))
}
