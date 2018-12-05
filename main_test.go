package gopool_test

import (
	"testing"
	"sync"
	"fmt"
	"time"
	"runtime"
	"github.com/binkzhao/gopool"
)

var taskCount = 1000000

type MyTask struct {
	Wg   *sync.WaitGroup
	Name string
}

func (mt *MyTask) Consume() error {
	mt.Wg.Done()
	time.Sleep(time.Second)
	return nil
}

func TestConsumeTaskWithPool(t *testing.T) {
	var wg sync.WaitGroup
	maxWorkers := 1000000
	maxQueues := 1000
	pool, _ := gopool.NewPool(maxWorkers, maxQueues)

	pool.Run()
	fmt.Println("Starting Worker Pool test...")

	startTime := time.Now()
	for i := 0; i < taskCount; i++ {
		wg.Add(1)
		myTask := &MyTask{
			Name: fmt.Sprintf("Test: %d", i),
			Wg:   &wg,
		}
		pool.Push(myTask)
	}

	wg.Wait()

	endTime := time.Now()
	diffTime := endTime.Sub(startTime) / time.Second
	fmt.Printf("[task count: %d]Use Go Pool All Task Spent %d second consume\n", taskCount, diffTime)
}

func TestNoPool(t *testing.T) {
	var wg sync.WaitGroup
	startTime := time.Now()
	for i := 0; i < taskCount; i++ {
		wg.Add(1)
		go func() {
			myTask := &MyTask{
				Name: fmt.Sprintf("Test: %d", i),
				Wg:   &wg,
			}
			myTask.Consume()
		}()
	}

	wg.Wait()

	endTime := time.Now()
	diffTime := endTime.Sub(startTime) / time.Second
	fmt.Printf("[task count: %d]No Use Pool All Task Spent %d second consume\n", taskCount, diffTime)
}
