package koi

import (
	"context"
	"log"

	"golang.org/x/sync/semaphore"
)

type Worker struct {
	QueueSize       int
	ConcurrentCount int64
	Work            func(any) any
}

type innerWorker struct {
	Worker
	ResultChan  chan any
	RequestChan chan any
	Semaphore   *semaphore.Weighted
}

func (i *innerWorker) job(request any) {
	defer i.Release()

	if result := i.Work(request); result != nil {
		i.ResultChan <- result
	}
}

func (i *innerWorker) Acquire() {
	err := i.Semaphore.Acquire(context.Background(), 1)
	if err != nil {
		log.Println("failed to acquire lock")
	}
}

func (i *innerWorker) Release() {
	i.Semaphore.Release(1)
}

func (w Worker) Validate() error {
	if w.ConcurrentCount < 1 {
		return errMinConcurrentCount
	}

	if w.QueueSize < 0 {
		return errNegativeQueueSize
	}

	return nil
}
