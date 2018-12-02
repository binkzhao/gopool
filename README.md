# Go pool
Goruntine Pool(简易版go协程池实现)，整体逻辑就是：先初始化Pool,并运行起来，把task放进任务队列taskQueue;
从队列取出任务task转发给Pool,从Pool中取出有效的worker进行处理task,处理完毕后再把worker重新丢进Pool里面
等待下次task调用。
支持逻辑：
- 可控制协程池大小，能批量处理大并发模式下的任务
- 任务类型为Task接口，支持业务逻辑自实现场景

# 任务类型
任务类型接口:
```
type Task interface {
	Consume() error
}
```
因此你要创建一个实现了该接口的类型，任何实现了该```task```的逻辑都可以使用该Pool.
例如：
```
type MyTask struct {
	Name string
}

func (mt *MyTask) Consume() error {
	fmt.Printf("[Task: %s] process done\n", mt.Name)
	return nil
}
```

## 如何使用
下载该代码包：
```
go get github.com/binkzhao/gocpool
```
使用默认的全局Pool
默认的全局Pool实现了单例全局唯一，默认maxWorkers=1000000,maxQueues=1000。
```
import (
	"github.com/binkzhao/gopool"
	"fmt"
	"time"
)

func main() {
	// start running pool
	gopool.Run()
	
	// push task to pool
	for i := 0; i <= 100; i++ {
		myTask := &MyTask{Name: fmt.Sprintf("Test: %d", i)}
		gopool.Push(myTask)
	}

	time.Sleep(time.Second * 30)
}
```
自定义pool size和queue size
```
func main() {
	maxWorkers := 100
	maxQueues := 100
	pool, _ := gopool.NewPool(maxWorkers, maxQueues)

	pool.Run()
	fmt.Println("Starting Worker Pool test...")

	for i := 0; i <= 100; i++ {
		myTask := &MyTask{Name: fmt.Sprintf("Test: %d", i)}
		pool.Push(myTask)
	}

	time.Sleep(time.Second * 30)
}
```
## 有待改进
- worker数据动态可控制，支持扩容和缩容。(pool定义workers切片来实现)
- 支持worker的有效期，worker可以根据有消费自动销毁(worker定义创建时间和pool设置一个全局的worker有效时间来实现)
- pool能够动态的停止退出等等。(设置退出通道信号量来实现)
