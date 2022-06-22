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

type Pond struct {
	Workers map[string]innerWorker
}

func NewPond() *Pond {
	return &Pond{
		Workers: make(map[string]innerWorker),
	}
}

func (p *Pond) RegisterWorker(id string, worker Worker) error {
	if err := worker.Validate(); err != nil {
		return err
	}

	innerWorker := innerWorker{
		Worker:      worker,
		ResultChan:  make(chan any, worker.QueueSize),
		RequestChan: make(chan any, worker.QueueSize),
		Semaphore:   semaphore.NewWeighted(worker.ConcurrentCount),
	}

	go p.manageWorker(innerWorker)

	p.Workers[id] = innerWorker

	return nil
}

func (p *Pond) AddJob(workerID string, request any) (chan any, error) {
	worker, ok := p.Workers[workerID]
	if !ok {
		return nil, errWorkerNotFound
	}

	// add request to worker queue
	worker.RequestChan <- request

	return worker.ResultChan, nil
}

func (p Pond) ResultChan(workerID string) chan any {
	worker, ok := p.Workers[workerID]

	if ok {
		return worker.ResultChan
	}

	return nil
}

func (p *Pond) manageWorker(worker innerWorker) {
	for request := range worker.RequestChan {
		worker.Acquire()

		go worker.job(request)
	}
}
