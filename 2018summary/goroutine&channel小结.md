
## 学习大纲
* [一.从并发模型说起](#1) 
* [二.goroutine的简介](#2) 
* [三.goroutine的使用姿势](#3) 
* [四.通道(channel)的简介](#4) 
* [五.重要的四种通道使用](#5) 
* [六.goroutine死锁与处理](#6) 
* [七.select的应用场景](#7) 
* [八.总结](#8) 



##  <span id="1">1.从并发模型说起</span>

并发目前比较主流的三种：
* 多线程
  每个线程一次处理一个请求，线程越多可并发处理的请求数就越多，但是在高并发下，多线程开销会比较大。
* 协程
  无需抢占式的调度，开销小，可以有效的提高线程的并发性，从而避免了线程的缺点的部分

* 异步回调IO模型
 说一个熟悉的，比如nginx使用的就是epoll模型，通过事件驱动的方式与异步IO回调，使得服务器持续运转，来支撑高并发的请求

 为了追求更高效和低开销的并发，golang的goroutine来了。


## <span id="2">2.goroutine的简介</span>
定义：在go里面，每一个并发执行的活动成为goroutine。
详解：goroutine可以认为是轻量级的线程，与创建线程相比，创建成本和开销都很小，每个goroutine的堆栈只有几kb，并且堆栈可根据程序的需要增长和缩小(线程的堆栈需指明和固定)，所以go程序从语言层面支持了高并发。

程序执行的背后：新的goroutine通过go关键字进行创建。

## <span id="3">3.goroutine 使用</span>


~~~go
package main

import (
	"fmt"
	"time"
)

func test() {
	fmt.Println(" a new  goroutine")
}

func main() {
	fmt.Println("start")
	go test()
	time.Sleep(1 * time.Second)
	fmt.Println("end")
}

~~~


## <span id="4">4.通道(channel)的简介</span>
~~~go
// ** nil 发送或接收都会堵塞
// A channel is in a nil state when it is declared to its zero value
var ch chan string

// A channel can be placed in a nil state by explicitly setting it to nil.
ch = nil

// ** open 允许发送、接收

// A channel is in a open state when it’s made using the built-in function make.
ch := make(chan string)

// ** closed 可以接收，不能发送

// A channel is in a closed state when it’s closed using the built-in function close.
close(ch)
~~~
## <span id="5">5.重要的四种通道使用</span>

[参考](http://39.106.173.209:88/channel-as-signalchang-jing-shi-yong-zong-jie/)
### 1.无缓冲--等待任务

~~~go
package main

import (
	"fmt"
)

func doLeaderWork() {
	fmt.Println("领导开始工作")
}
func doEmpWork() {
	fmt.Println("员工开始工作")
}

func main() {
	fmt.Println("start")
	var ch = make(chan struct{})

	// goroutine 等待领导任务
	go func() {

		doEmpWork()
		<-ch //堵塞于此，等着 send 后执行【接收发送端数据，在发送之前，堵塞于此】
	}()
	doLeaderWork()

	ch <- struct{}{} //领导做完后，发送信号，员工再做

	fmt.Println("end")
}

~~~
### 2.无缓冲--等待结果

~~~go
package main

import (
	"fmt"
)

func doLeaderWork() {
	fmt.Println("领导开始工作")
}
func doEmpWork() {
	fmt.Println("员工开始工作")
}

func main() {
	fmt.Println("start")
	var ch = make(chan struct{})
   //员工先做，领导等待员工的结果再做
	go func() {

		ch <- struct{}{}
		doEmpWork()
	}()

	<-ch //堵塞于此，等待员工做完后，领导再做。【接收发送在发送之前，堵塞于此】
	doLeaderWork()

	fmt.Println("end")
}


~~~

### 3.有缓冲 无保证
~~~go
package main

import (
	"fmt"
)

func doLeaderWork() {
	fmt.Println("领导开始工作")
}
func doEmpWork() {
	fmt.Println("员工开始工作")
}

func main() {
	fmt.Println("start")
	emp := 5
	var ch = make(chan struct{}, emp)

	for i := 0; i < emp; i++ {
		go func(who int) {
			doEmpWork()
			ch <- struct{}{}
		}(i)

	}

	for emp > 0 {
		<-ch //堵塞于此，等待员工做完,领导再做
		doLeaderWork()
		emp--
		fmt.Println("............")
	}
	fmt.Println("end")
}

~~~

~~~go
package main

import (
	"fmt"
	"time"
)

func doLeaderWork() {
	fmt.Println("领导开始工作")
}
func doEmpWork() {
	fmt.Println("员工开始工作")
}

func main() {
	fmt.Println("start")
	emp := 5
	var ch = make(chan struct{}, emp)

	go func() {
		for _ = range ch {

			doEmpWork()
		}

	}()
	const work = 8
	for i := 0; i < work; i++ {
		select {
		case ch <- struct{}{}:
			doLeaderWork()
		default:
			fmt.Println("容量慢了，不能再发送了")

		}
	}
	time.Sleep(1e9)
	close(ch) //关闭channel，发送结束的信号给员工
	fmt.Println("end")
}

~~~

### 4.有缓冲，延迟保证
~~~go
package main

import (
	"fmt"
)

func doLeaderWork() {
	fmt.Println("领导开始工作")
}
func doEmpWork() {
	fmt.Println("员工开始工作")
}

func main() {
	fmt.Println("start")
	emp := 1
	var ch = make(chan struct{}, emp)

	go func() {
		for _ = range ch {

			doEmpWork()

		}
	}()
	for i := 0; i < 8; i++ {
		doLeaderWork()
		ch <- struct{}{}

	}
	close(ch)

	fmt.Println("end")
}

~~~
## <span id="6">6.goroutine死锁与处理</span>
### 1.死锁场景一
~~~go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	ch := make(chan struct{})
	<-ch//无缓冲，只有接收，没有发生，将main goroutine堵塞于此

	fmt.Println("end")
}

~~~
输出：
~~~
start
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
~~~

### 2.死锁场景二
~~~go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	ch := make(chan struct{})
	ch2 := make(chan struct{})
	go func() {
        fmt.Println("a new goroutine")
		ch <- struct{}{}//未被接收，堵塞与此
		ch2 <- struct{}{}
		
	}()
	<-ch2

	fmt.Println("end")
}

~~~


输出：
~~~
start
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
~~~

### 3.死锁原因及处理
* 原因：无缓冲channel，有发送就一定对应接收，两者缺少一个必定造成堵塞，发送死锁。
* 处理：将没有接收或者没有发生的操作完成即可或者将无缓冲的改成有缓冲的。

   - 场景一：
   
       ~~~go
       package main

        import (
            "fmt"
        )

        func main() {
            fmt.Println("start")
            ch := make(chan struct{})

            go func() {
                ch <- struct{}{}
                fmt.Println("a new goroutine")

            }()
            <-ch

            fmt.Println("end")
       }

       ~~~
   - 场景二---有发送与接收对应
  

     ~~~go
     package main

        import (
            "fmt"
        )

        func main() {
            fmt.Println("start")
            ch := make(chan struct{})
            ch2 := make(chan struct{})
            go func() {
                fmt.Println("a new goroutine")
                ch <- struct{}{} 
                ch2 <- struct{}{}

            }()
            <-ch
            <-ch2
            //time.Sleep(time.Second) //暂定一下，看一下输出结果

            fmt.Println("end")
        }

     ~~~
   -  场景二---有缓冲处理
    
     
     ~~~go
     
        package main

        import (
            "fmt"
        )

        func main() {
            fmt.Println("start")
            ch := make(chan struct{}, 1) //有缓冲，放入1个不会堵塞
            ch2 := make(chan struct{})
            go func() {
                fmt.Println("a new goroutine")
                ch <- struct{}{}
                ch2 <- struct{}{}

            }()

            <-ch2

            fmt.Println("end")
        }

     ~~~
## <span id="7">7.select的应用场景</span>

### 1.超时作用

~~~go
package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	select {
	case <-ch:
	case <-time.After(time.Second):
		fmt.Println("超时退出!")
	}

	fmt.Println("end")

}
~~~
### 2.退出作用

### 3.判断channel是否已满(是否堵塞)



## <span id="8">8.总结</span>
* Goroutine: 让研发人员更加专注于业务逻辑，从os层面的逻辑抽离出来。
* Channel:简单、安全、高效的实现了多个goroutine之间的同步与信息传递。
* Select:可以处理多个channel。
> reference:

[csdn](https://blog.csdn.net/u011957758/article/details/81159481)

