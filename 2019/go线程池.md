### 本节大纲
* [一、背景](#1)
* [二、go线程池模型](#2)
* [三、设计方案(考虑的问题)](#3)
* [四、pool源码解析](#4)
* [五、调用实例](#5)


### <span id="1">一、背景</span>
使用goroutine用户态线程，可以避免os调度器切换，两级线程模型M:N,可以充分复用内核线程；另外简单粗暴开启go func不是最优方案，容易造成内存暴涨、gc频繁。
### <span id="2">二、go线程池模型</span>
  生产者--->队列(channel)--->消费者模型。
### <span id="3">三、设计方案(考虑的问题)</span>
* 1.限制goroutine数量及重用,否则大量goroutine造成内存暴涨、gc频繁(cpu只高不下) 
* 2.任务队列的设计 
* 3.优雅关闭线程池，保证所有任务都执行完毕


### <span id="4">四、pool源码解析</span>
~~~

package pool

import (
	"errors"
	"fmt"
	"log"
	"sync/atomic"
)


var (
	ErrInvalidPoolCnt = errors.New("invalid pool size cnt")
	ErrPoolHasClosed  = errors.New("pool has closed")
)

const (
	RUNNING = 1

	STOPED = 2
)

// TaskHandler 需要执行任务的封装
type TaskHandler struct {
	Handler func(v ...interface{}) error
	Params  []interface{}
}

// Pool 任务池，基于生产者--->队列--->消费者模式
//
type Pool struct {
	sizeCnt           uint64            //任务池总大小
	currentRunningCnt uint64            //当前运行goroutine数量
	status            int64             //任务池状态(运行、关闭)
	taskC             chan *TaskHandler //任务队列channel
	closeC            chan bool         //关闭channel信号
	PanicHandler      func(interface{}) //针对每一个task 相应进行panic处理
}

func NewPool(sizeCnt uint64) (*Pool, error) {
	if sizeCnt <= 0 {
		return nil, ErrInvalidPoolCnt
	}
	return &Pool{
		sizeCnt: sizeCnt,
		status:  RUNNING,
		taskC:   make(chan *TaskHandler, sizeCnt), //初始化任务队列，默认最大长度
		closeC:  make(chan bool),
	}, nil
}

// GetSizeCnt 获取线程池当前容量大小
func (p *Pool) GetSizeCnt() uint64 {
	return atomic.LoadUint64(&p.sizeCnt)
}

// GetCurrentRunningCnt 获取正在运行goroutine数量
func (p *Pool) GetCurrentRunningCnt() uint64 {
	return atomic.LoadUint64(&p.currentRunningCnt)
}

//incCurrentRunningCnt 新增任务时，增加goroutine数量
func (p *Pool) incCurrentRunningCnt() {
	atomic.AddUint64(&p.currentRunningCnt, 1)
}

//decCurrentRunningCnt 某个任务执行结束后，减少goroutine数量
func (p *Pool) decCurrentRunningCnt() {
	atomic.AddUint64(&p.currentRunningCnt, ^uint64(0))
}

// 新增任务：如果pool满了，则放在队列中
func (p *Pool) AddTask(task *TaskHandler) error {

	if p.status == STOPED {
		return ErrPoolHasClosed
	}

	//如果当前运行goroutine数量池中不满，则直接运行，如果满了，则直接放在队列channel中
	if p.GetCurrentRunningCnt() < p.GetSizeCnt() {
		p.run()
	}

	p.taskC <- task

	return nil
}

func (p *Pool) run() {
	p.incCurrentRunningCnt()

	go func() {
		defer func() {
			p.decCurrentRunningCnt()
			if r := recover(); r != nil {
				if p.PanicHandler != nil {
					p.PanicHandler(r)
				} else {
					log.Printf("task handler panic: %s\n", r)
				}
			}
		}()

		for {
			select {
			case task, ok := <-p.taskC: //从任务队列channel中获取任务
				if !ok {
					return //如果任务队列channel关闭了，则返回
				}
				fmt.Println("task", task)
				task.Handler(task.Params...) //处理任务
			case <-p.closeC: //收到关闭pool的信号
				return
			}
		}
	}()
}

// Close 关闭线程池
func (p *Pool) Close() {
	p.status = STOPED //设置关闭状态，不接收新任务

	//等待所以task消费完成
	for len(p.taskC) > 0 {
	}
	//发送关闭pool的信号
	p.closeC <- true
	//生产端关闭任务队列channel
	close(p.taskC)
}

~~~

### <span id="5">五、调用实例</span>
~~~

func doHandler(p ...interface{}) error {
	fmt.Println("p:", p)
	return nil
}
func main() {

	wg := sync.WaitGroup{}

	// 创建任务池,限制goroutine数量并重用，sizeCnt根据业务调整其大小
	p, err := pool.NewPool(5)
	if err != nil {
		panic(err)
	}

	//taskHandler := &pool.TaskHandler{
	//	Handler: doHandler,
	//	//Params:  []interface{}{1},
	//}

	arrTask := []int{1, 3, 5, 7, 9, 11, 13, 15}

	for i := 0; i < len(arrTask); i++ {
		wg.Add(1)
		// 任务放入池中
		//taskHandler.Params = []interface{}{arrTask[i]}
		//		//p.AddTask(taskHandler)

		// 任务放入池中
		p.AddTask(&pool.TaskHandler{
			Handler: func(v ...interface{}) error {
				defer wg.Done()
				return doHandler(v)
			},
			Params: []interface{}{arrTask[i]},
		})

	}

	wg.Wait() //等待执行
	//time.Sleep(2 * time.Second) // 等待执行
}
~~~
