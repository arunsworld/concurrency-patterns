package main

import (
	"context"
	"log"
	"time"

	concurrencypatterns "github.com/arunsworld/concurrency-patterns/concurrency-patterns"
	"github.com/arunsworld/nursery"
)

func main() {
	result := &concurrencypatterns.Job{}
	count := 100
	ch := make(chan concurrencypatterns.Job, 10)
	start := time.Now()
	nursery.RunConcurrently(
		// the 100 independent jobs
		func(context.Context, chan error) {
			worker := concurrencypatterns.NewWorker(concurrencypatterns.SetAveDelayOnWorker(1000), concurrencypatterns.SetStdDeviationOnWorker(300))
			nursery.RunMultipleCopiesConcurrently(count,
				func(ctx context.Context, _ chan error) {
					jobID := ctx.Value(nursery.JobID).(int)
					worker.Work(ctx, false) // do the actual work
					job := &concurrencypatterns.Job{}
					job.MarkIntermediateState(jobID)
					ch <- *job
				},
			)
			close(ch)
		},
		// bringing it altogether
		func(ctx context.Context, _ chan error) {
			worker := concurrencypatterns.NewWorker(concurrencypatterns.SetAveDelayOnWorker(10), concurrencypatterns.SetStdDeviationOnWorker(0))
			for j := range ch {
				worker.Work(ctx, false)
				result.Merge(j)
			}
		},
	)
	if !result.Verify(count) {
		log.Fatalf("verification failed: %v", time.Since(start))
	}
	log.Printf("finished in: %v", time.Since(start))
}
