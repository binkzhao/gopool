package gopool

import "sync"

const (
	DefaultMaxQueues  = 100
	DefaultMaxWorkers = 10000
)

// default pool
var DefaultPool *pool
var once sync.Once

func init() {
	once.Do(func() {
		DefaultPool, _ = NewPool(DefaultMaxWorkers, DefaultMaxQueues)
	})
}

// submit task to worker pool
func Submit(task Task) {
	DefaultPool.PushTask(task)
}

// Running default goroutines pool
func Running() {
	DefaultPool.Run()
}

// Closes the default pool.
func Stop() {
	DefaultPool.Stop()
}
