package gopool

import (
	"log"
)

// Worker represents the worker that executes the task
type worker struct {
	workerPool  chan chan Task // worker pool
	taskChannel chan Task      // task chan
	quit        chan bool
}

func NewWorker(workerPool chan chan Task) worker {
	return worker{
		workerPool:  workerPool,
		taskChannel: make(chan Task),
		quit:        make(chan bool),
	}
}

func (w *worker) Start() {
	go func() {
		for {
			// consume done ,then worker reenter workerPool
			w.workerPool <- w.taskChannel
			select {
			case task := <-w.taskChannel:
				// received a work request and consume it
				if err := task.Consume(); err != nil {
					log.Printf("Consume fail: %v", err.Error())
				}
			case <-w.quit:
				return
			}
		}
	}()
}

// TODO
func (w *worker) Stop() {
	go func() {
		w.quit <- true
	}()
}
