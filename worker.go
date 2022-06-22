package koi

import (
	"context"
	"log"

	"golang.org/x/sync/semaphore"
)

type Worker[T any, E any] struct {
	QueueSize       int
	ConcurrentCount int64
	Work            func(T) *E
}

type innerWorker[T any, E any] struct {
	Worker      Worker[T, E]
	ResultChan  chan *E
	RequestChan chan T
	Semaphore   *semaphore.Weighted
}

func (i *innerWorker[T, E]) work(request T) {
	defer i.Release()

	if result := i.Worker.Work(request); result != nil {
		i.ResultChan <- result
	}
}

func (i *innerWorker[T, E]) Acquire() {
	err := i.Semaphore.Acquire(context.Background(), 1)
	if err != nil {
		log.Println("failed to acquire lock")
	}
}

func (i *innerWorker[T, E]) Release() {
	i.Semaphore.Release(1)
}

func (w Worker[T, E]) Validate() error {
	if w.ConcurrentCount < 1 {
		return errMinConcurrentCount
	}

	if w.QueueSize < 0 {
		return errNegativeQueueSize
	}

	return nil
}
