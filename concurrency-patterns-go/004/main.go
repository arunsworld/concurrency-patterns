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
	result := make(chan string)
	stopCtx, stop := context.WithCancel(context.Background())
	nursery.RunConcurrently(
		func(context.Context, chan error) {
			defer stop()
			worker := concurrencypatterns.NewWorker(concurrencypatterns.SetStdDeviationOnWorker(0))
			worker.Work(stopCtx, false)
			if !nursery.IsContextDone(stopCtx) {
				result <- "job 1"
			}
			log.Printf("job 1 finished in: %v", time.Since(start))
		},
		func(context.Context, chan error) {
			defer stop()
			worker := concurrencypatterns.NewWorker(concurrencypatterns.SetAveDelayOnWorker(1000), concurrencypatterns.SetStdDeviationOnWorker(0))
			worker.Work(stopCtx, false)
			if !nursery.IsContextDone(stopCtx) {
				result <- "job 2"
			}
			log.Printf("job 2 finished in: %v", time.Since(start))
		},
		func(context.Context, chan error) {
			log.Printf("received result from: %s", <-result)
		},
	)
}
