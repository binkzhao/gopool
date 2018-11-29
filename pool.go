package gopool

type pool struct {
	workerPool chan chan Task // worker pool
	taskQueue  chan Task      // work queue
	maxWorkers int            // worker pool max worker count
	curWorkers int            // current worker count
	quit       chan bool
}

func NewPool(maxWorkers, maxQueues int) (*pool, error) {
	p := &pool{
		workerPool: make(chan chan Task, maxWorkers),
		taskQueue:  make(chan Task, maxQueues),
		maxWorkers: maxWorkers,
	}

	return p, nil
}

func (p *pool) Run() {
	for i := 0; i < p.maxWorkers; i++ {
		w := NewWorker(p.workerPool)
		w.Start()
	}
	go p.dispatch()
}

// submit task to Task queue
func (p *pool) PushTask(task Task) error {
	p.taskQueue <- task
	return nil
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
		case <-p.quit:
			return
		}
	}
}

// TODO
func (p *pool) Stop() {
	go func() {
		p.quit <- true
	}()
}
