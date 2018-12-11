
package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

var ErrTimeOut = errors.New("执行超时")
var ErrInterrupt = errors.New("执行中断")

//一个worker,包含多个要执行的tasks
type Worker struct {
	tasks     []func(int)      //要执行的任务
	complete  chan error       //用于通知任务全部完成
	timeout   <-chan time.Time //任务在多久内完成
	interrupt chan os.Signal   //强制终止的信号

}

func NewWork(tm time.Duration) *Worker {
	return &Worker{
		complete:  make(chan error),
		timeout:   time.After(tm),
		interrupt: make(chan os.Signal, 1),
	}
}

//将需要执行的任务，添加到Worker里
func (r *Worker) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

//执行的过程中如何接收到中断信号时，则返回中断错误
//如果任务全部执行完，还没有接收到中断信号，则返回nil
func (r *Worker) RunWorker() error {
	for id, task := range r.tasks {
		if r.IsInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

//是否接收到了中断信号
func (r *Worker) IsInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

//开始执行所有任务，并且监视通道事件
func (r *Worker) Start() error {
	//希望接收哪些系统信号
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.RunWorker()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOut
	}
}

func main() {
	log.Println("start task")

	//设置超时时间为3秒，添加4个生成的任务===>超时
	//设置超时时间 5秒，添加4个任务===》任务执行结束
	//设置超时时间 秒，添加4个任务,执行过程中ctrl+c ===》执行中断
	timeout := 10 * time.Second
	r := NewWork(timeout)

	r.Add(createTask(), createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case ErrTimeOut:
			log.Println(err)
			os.Exit(1)
		case ErrInterrupt:
			log.Println(err)
			os.Exit(2)
		}
	}
	log.Println("end task")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("正在执行任务第%d任务", id)
		//模拟每个task执行1s
		time.Sleep(2 * time.Second)
	}
}

//参考：http://www.flysnow.org/2017/04/29/go-in-action-go-runner.html
