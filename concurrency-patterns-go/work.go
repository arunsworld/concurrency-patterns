package concurrencypatterns

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func NewWorker(opts ...WorkerOption) worker {
	w := worker{
		aveDelayMs:     100,
		stdDeviationMs: 10,
	}
	for _, opt := range opts {
		w = opt(w)
	}
	return w
}

type worker struct {
	aveDelayMs     int
	stdDeviationMs int
}

func (w worker) Work(ctx context.Context, returnErr bool) error {
	delay := rand.NormFloat64()*float64(w.stdDeviationMs) + float64(w.aveDelayMs)
	if delay < 0 {
		delay = 0
	}
	select {
	case <-time.After(time.Millisecond * time.Duration(delay)):
	case <-ctx.Done():
	}
	return w.err(returnErr)
}

func (w worker) err(returnErr bool) error {
	if !returnErr {
		return nil
	}
	return fmt.Errorf("worker resulted in an error")
}

type WorkerOption func(w worker) worker

var SetAveDelayOnWorker = func(v int) WorkerOption {
	return func(w worker) worker {
		w.aveDelayMs = v
		return w
	}
}

var SetStdDeviationOnWorker = func(v int) WorkerOption {
	return func(w worker) worker {
		w.stdDeviationMs = v
		return w
	}
}
