### 背景

  定时任务，每隔3s执行一次 WithdrawTransfer中的DoTask()。


### 目录

~~~
tom@tomdeMacBook-Pro:~/go/src/demo$ tree
.
├── main.go
└── task
    └── transfer.go
~~~

### main.go
~~~
package main

import (
	"demo/task"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/net/context"
)

var transfer *task.WithdrawTransfer

// func init() {

// 	transfer = task.NewWithdrawTransfer(context.Background(), time.Second*3)
// }

func main() {
	defer func() {
		time.Sleep(time.Second * 3)
		if r := recover(); r != nil {
			log.Printf("[main] recover painc:%s.", r)
			exit(-1)
		}
		//log.Printf("main quited safely, saved %d songs\r\n", changba.Cb.Count)
	}()

	go func() {
		signals := make(chan os.Signal)
		signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
		<-signals
		exit(0)
	}()

	log.Println("task started")
	transfer = task.NewWithdrawTransfer(context.Background(), time.Second*3)
	transfer.Run()
	log.Println("main end ")
}

func exit(status int) {
	log.Printf("[main] exit:%d.", status)
	if status == 0 {
		transfer.Stop()
	}
	os.Exit(status)
}


~~~


### task/transfer.go

~~~
package task

import (
	"fmt"
	"time"

	"golang.org/x/net/context"
)

type WithdrawTransfer struct {
	ctx    context.Context
	cancel context.CancelFunc
	ticker *time.Ticker
}

func NewWithdrawTransfer(ctx context.Context, interval time.Duration) *WithdrawTransfer {
	ctx, cancel := context.WithCancel(ctx)
	ticker := time.NewTicker(interval)
	return &WithdrawTransfer{
		ctx:    ctx,
		cancel: cancel,
		ticker: ticker,
	}

}

func (w *WithdrawTransfer) Stop() {
	w.cancel()
	w.ticker.Stop()
}
func (w *WithdrawTransfer) DoTask() {
	fmt.Println("do DoTask  task....")
}

func (w *WithdrawTransfer) Run() {
  for {
		select {
		case <-w.ticker.C:

			w.DoTask()

		case <-w.ctx.Done():
			fmt.Println("cancel func called, transfer will exited now")
			return
		}
	}
}

~~~
