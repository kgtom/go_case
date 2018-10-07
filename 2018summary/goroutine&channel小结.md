
## 学习大纲
* [1.从并发模型说起](#1) 
* [2.goroutine的简介](#2) 
* [3.goroutine的使用姿势](#3) 
* [4.通道(channel)的简介](#4) 
* [5.重要的四种通道使用](#5) 
* [6.goroutine死锁与处理](#6) 
* [7.select的简介](#7) 
* [8.select的应用场景](#8) 
* [9.select死锁](#9)


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


## <span id="4">4.goroutine 使用</span>

> reference:

[csdn](https://blog.csdn.net/u011957758/article/details/81159481)

