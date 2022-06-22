package koi

import (
	"errors"

	"golang.org/x/sync/semaphore"
)

var (
	errWorkerNotFound     = errors.New("worker Not Found")
	errMinConcurrentCount = errors.New("concurrent count must be at least 1")
	errNegativeQueueSize  = errors.New("worker request queue size can not be negative")
)

type Pond[T any, E any] struct {
	Workers map[string]innerWorker[T, E]
}

func NewPond[T any, E any]() *Pond[T, E] {
	return &Pond[T, E]{
		Workers: make(map[string]innerWorker[T, E]),
	}
}

func (p *Pond[T, E]) RegisterWorker(id string, worker Worker[T, E]) error {
	if err := worker.Validate(); err != nil {
		return err
	}

	innerWorker := innerWorker[T, E]{
		Worker:      worker,
		ResultChan:  make(chan *E, worker.QueueSize),
		RequestChan: make(chan T, worker.QueueSize),
		Semaphore:   semaphore.NewWeighted(worker.ConcurrentCount),
	}

	go p.manageWorker(innerWorker)

	p.Workers[id] = innerWorker

	return nil
}

func (p *Pond[T, E]) AddWork(workerID string, request T) (chan *E, error) {
	worker, ok := p.Workers[workerID]
	if !ok {
		return nil, errWorkerNotFound
	}

	// add request to worker queue
	worker.RequestChan <- request

	return worker.ResultChan, nil
}

func (p Pond[T, E]) ResultChan(workerID string) chan *E {
	worker, ok := p.Workers[workerID]

	if ok {
		return worker.ResultChan
	}

	return nil
}

func (p *Pond[T, E]) manageWorker(worker innerWorker[T, E]) {
	for request := range worker.RequestChan {
		worker.Acquire()

		go worker.work(request)
	}
}
