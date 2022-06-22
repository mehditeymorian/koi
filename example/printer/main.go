package main

import (
	"log"
	"time"

	"github.com/mehditeymorian/koi/v2"
)

func main() {
	pond := koi.NewPond[int, int]()

	printWorker := koi.Worker[int, int]{
		ConcurrentCount: 2,
		QueueSize:       10,
		Work: func(a int) *int {
			time.Sleep(1 * time.Second)
			log.Println(a)

			return nil
		},
	}

	_ = pond.RegisterWorker("printer", printWorker)

	for i := 0; i < 10; i++ {
		_, err := pond.AddWork("printer", i)
		if err != nil {
			log.Printf("error while adding job: %v\n", err)
		}
	}

	log.Println("all job added")

	for {

	}
}
