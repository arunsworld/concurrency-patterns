package main

import (
	"context"
	"log"
	"time"

	concurrencypatterns "github.com/arunsworld/concurrency-patterns/concurrency-patterns"
	"github.com/arunsworld/nursery"
)

func main() {
	worker := concurrencypatterns.NewWorker()

	start := time.Now()
	err := nursery.RunMultipleCopiesConcurrently(100,
		func(ctx context.Context, errCh chan error) {
			if err := worker.Work(ctx, false); err != nil {
				errCh <- err
			}
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("finished in: %v", time.Since(start))
}
