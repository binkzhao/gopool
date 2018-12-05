package gopool

import (
	"fmt"
	"errors"
)

// pool struct
type pool struct {
	workerPool chan chan Task // worker pool
	taskQueue  chan Task      // work queue
	maxWorkers int            // worker pool max worker count
}

// create a pool object
func NewPool(maxWorkers, maxQueues int) (*pool, error) {
	p := &pool{
		workerPool: make(chan chan Task, maxWorkers),
		taskQueue:  make(chan Task, maxQueues),
		maxWorkers: maxWorkers,
	}

	return p, nil
}

// run this pool
func (p *pool) Run() {
	p.initWorkers()
	go p.dispatch()
}

// submit task to Task queue
func (p *pool) Push(task Task) error {
	var err error
	defer func() {
		// handle taskQueue channel maybe closed
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("Push Task Fail: %v", e))
		}
	}()

	p.taskQueue <- task
	return err
}

// init pool workers
func (p *pool) initWorkers() {
	for i := 0; i < p.maxWorkers; i++ {
		w := NewWorker(p.workerPool)
		w.Start()
	}
}

// dispatch task
func (p *pool) dispatch() {
	for {
		select {
		case task := <-p.taskQueue:
			// a task request has been received
			go func(task Task) {
				// pop a task channel from pool
				// this will block until a worker is idle
				taskChannel := <-p.workerPool
				// push task to task channel
				taskChannel <- task
			}(task)
		}
	}
}
