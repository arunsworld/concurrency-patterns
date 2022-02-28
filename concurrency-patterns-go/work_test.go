package concurrencypatterns_test

import (
	"context"
	"log"
	"testing"
	"time"

	concurrencypatterns "github.com/arunsworld/concurrency-patterns/concurrency-patterns"
)

func TestWorker(t *testing.T) {
	w := concurrencypatterns.NewWorker()
	for i := 0; i < 100; i++ {
		s := time.Now()
		w.Work(context.Background(), false)
		log.Printf("delay: %v", time.Since(s))
	}
}
