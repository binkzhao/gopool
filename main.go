package gopool

import "sync"

const (
	// default task queue max size
	DefaultMaxQueues = 1000
	// default pool worker max size
	DefaultMaxWorkers = 1000000
)

var (
	// default pool
	DefaultPool *pool
	// control default pool init
	once sync.Once
)

// use sync.once init Default pool
func init() {
	once.Do(func() {
		DefaultPool, _ = NewPool(DefaultMaxWorkers, DefaultMaxQueues)
	})
}

// submit task to worker pool
func Push(task Task) error {
	err := DefaultPool.Push(task)
	return err
}

// Running default goroutines pool
func Run() {
	DefaultPool.Run()
}
