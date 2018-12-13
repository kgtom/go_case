## 问题索引
 * [一.接收消息队列中消息，并发处理与批处理](#1)
 * [二.超时控制context.WithDeadline](#2)
 * [三.性能分析--trace/pprof](#3)
 * [四.error处理](#4)
 * [五.何时选择并发](#5)
 * [六.map线程不安全](#6)
 * [七.组合优于继承](#7)


### <span id="1">一.接收消息队列中消息，并发处理与批处理</span>


**问题：** 
~~~go
func main() {

	
	fmt.Println("start")
	ch := make(chan *Message, 20)
	for i := 0; i <= 12; i++ {

		m := &Message{Nums: i}
		ch <- m
		go a(ch)
	}
	//暂停看一下结果
	time.Sleep(1 * time.Second)
	fmt.Println("end")
}

type Message struct {
	Nums int
}

//并发少的情况下没有问题，一旦并发上来，请求堆积，造成雪崩。
func a(nsqMsg <-chan *Message) {

	for {
		select {
		case msg := <-nsqMsg:
			handleA(msg)
		}

	}
}

//只读取 channel  (m chan <-*Message)
func handleA(m *Message) {

	fmt.Println("msg:", *m)

}
~~~


**解决方案：** 一种并发处理、另一种 批量处理
~~~go
func main() {

	
	fmt.Println("start")
	ch := make(chan *Message, 20)
	for i := 0; i <= 12; i++ {

		m := &Message{Nums: i}
		ch <- m
		//go a(ch)
		go b(ch)

	}
	//暂停看一下结果
	time.Sleep(1 * time.Second)
	fmt.Println("end")
}

type Message struct {
	Nums int
}

//并发处理 接收到nsq 消息，使用缓冲区现在并发规模
func a(nsqMsg chan *Message) {

	m := make(chan *Message, 10)
	for {
		select {
		case msg := <-nsqMsg:
			m <- msg
			go func(m <-chan *Message) {
				handleA(m)
			}(m)
		}

	}
}

//只读取 channel  (m chan <-*Message)
func handleA(m <-chan *Message) {

	for v := range m {
		fmt.Println("msg:", *v)
	}
}

//批量处理 接收nsq信息，一并处理
func b(nsqMsg chan *Message) {

	for {
		select {
		case msg := <-nsqMsg:
			messages := []*Message{msg}
			for i := 0; i < len(nsqMsg); i++ {
				messages = append(messages, <-nsqMsg)
			}
			handleB(messages)
		}
	}
}
func handleB(msg []*Message) {

	for k, v := range msg {
		fmt.Println("num:", k, "msg：", v)
	}
}

~~~

### 二.超时控制context.WithDeadline

* 调用外部请求(http、rpc)时或者核心代码，必须要有超时控制，做到高可用中的隔离。
~~~go

func main() {

	//超时控制,3秒没有请求返回则返回失败
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	defer cancel()
	select {

	case ret := <-Handle():
		fmt.Println("ok ", ret)
	case <-ctx.Done():
		fmt.Println("ctx:", ctx.Err())

	}
}
func Handle() <-chan struct{} {

	ch := make(chan struct{})
	go func() {
		//do something
		time.Sleep(4 * time.Second)
		ch <- struct{}{}
	}()

	return ch
}

~~~

### <span id="3">三.性能调优-- trace/pporf</span>
#### trace 跟踪器
 - trace 作用：监听 Go 运行时的一些特定的事件，如：
1. goroutine 的创建、开始和结束。
2. 阻塞/解锁goroutine的一些事件（系统调用，channel，锁）
3. 网络I/O、系统调用、GC
- 解决我们遇到问题：
* goroutine 如何被调度的？什么时候被堵塞？
* 为什么有时候并发的程序反而不如串行执行快？

###### 1.非web --"runtime/trace"

~~~go
package main

import (
	"os"
    "fmt"
	"runtime/trace"
)

func main() {
	traceTest()
}

func traceTest() {

	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	//your code
	ch := make(chan struct{})

	go func() {

		ch <- struct{}{}
	}()

	<-ch
   fmt.Println("end")
}

~~~
* 使用 go tool trace trace.out 自动打开web页面
~~~
# tom @ tom-pc in ~/goprojects/src/goroutine-demo [21:20:02]
$ go tool trace trace.out
2018/12/13 21:20:07 Parsing trace...
2018/12/13 21:20:07 Serializing trace...
2018/12/13 21:20:07 Splitting trace...
2018/12/13 21:20:07 Opening browser. Trace viewer is listening on http://127.0.0.1:59900

~~~
* 打开可视化界面，进行跟踪器查询、goroutine 跟踪及堆栈信息等。

###### 2.web 跟踪 --"_net/http/pprof"

* 代码如下
~~~go
package main

import (
	"net/http"

	"fmt"
	_ "net/http/pprof"
	"sync"
)

func main() {
	http.HandleFunc("/test", test)

	http.ListenAndServe(":8081", nil)
}

func test(w http.ResponseWriter, r *http.Request) {

	process()
	fmt.Fprint(w, "success")
}

func doSomething(i int) {
	fmt.Println("do:", (i+2)*2/2)
}

func process() (int, error) {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			doSomething(i)

		}(i)
	}
	wg.Wait()

	return http.StatusOK, nil
}

~~~
* 发送请求，收集跟踪信息
~~~

# tom @ tom-pc in ~ [22:08:42] C:1
$ curl localhost:8081/debug/pprof/trace?seconds=10 > trace.out
zsh: no matches found: localhost:8081/debug/pprof/trace?seconds=10
~~~

* 遇到 zsh:no matches found 的问题，解决：在~/.zshrc中加入:setopt no_nomatch即可。 
~~~
# tom @ tom-pc in ~ [22:05:21] C:1
$ vi ~/.zshrc

...
setopt no_nomatch
...

# tom @ tom-pc in ~ [22:05:55]
$ source ~/.zshrc

~~~
* 再次请求并打开

~~~

# tom @ tom-pc in ~ [22:18:22] C:130
$ curl localhost:8081/debug/pprof/trace?seconds=10 > trace.out
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100  2490    0  2490    0     0    248      0 --:--:--  0:00:10 --:--:--   520

# tom @ tom-pc in ~ [22:18:41]
$ go tool trace trace.out
2018/12/13 22:19:10 Parsing trace...
2018/12/13 22:19:10 Serializing trace...
2018/12/13 22:19:10 Splitting trace...
2018/12/13 22:19:10 Opening browser. Trace viewer is listening on http://127.0.0.1:50367


~~~
- 也可以使用wrk 压测的同时，在另一个终端使用 curl localhost:8081/debug/pprof/trace?seconds=10 > trace.out，然后再用 go tool trace trace.out 命令查看。
~~~
# tom @ tom-pc in ~ [22:20:02] C:130
$ wrk -c 100 -t 10 -d 30s http://localhost:8081/test
Running 30s test @ http://localhost:8081/test
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    47.09ms   32.45ms 400.11ms   83.44%
    Req/Sec   229.04     78.73   460.00     65.13%
  68499 requests in 30.09s, 8.04MB read
Requests/sec:   2276.54
Transfer/sec:    273.45KB
~~~

#### pporf --内存和cpu利用情况
 ###### 1.web应用程序:http://localhost:8081/test

~~~go
package main

import (
	"log"
	"net/http"

	"fmt"
	_ "net/http/pprof"
	"sync"
)

func main() {
	http.HandleFunc("/test", test)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func test(w http.ResponseWriter, r *http.Request) {

	process()
	fmt.Fprint(w, "success")
}

func doSomething(i int) {
	fmt.Println("do:", i)
}

func process() (int, error) {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			doSomething(i)

		}(i)
	}
	wg.Wait()

	return http.StatusOK, nil
}


~~~
* 执行压测： 30s内4个线程 200个链接请求

~~~
# tom @ tom-pc in ~ [22:18:06]
$ wrk -c 200 -t 4 -d 30s http://localhost:8081/test
Running 30s test @ http://localhost:8081/test
  4 threads and 200 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    18.65ms   18.87ms 321.55ms   93.51%
    Req/Sec     3.09k   796.34     5.51k    69.46%
  368907 requests in 30.10s, 43.27MB read
Requests/sec:  12257.01
Transfer/sec:      1.44MB
~~~
* 浏览器打开：查看cpu： http://localhost:8081/debug/pprof/profile 稍后片刻，可以下载到文件 profile(mac 默认下载地址： ```~/Downloads/profile```)。(查看内存：http://localhost:8081/debug/pprof/heap),以查看cpu为例：
~~~
# tom @ tom-pc in ~ [22:19:48]
$ go tool pprof ~/Downloads/profile
Main binary filename not available.
Type: cpu
Time: Dec 13, 2018 at 22:16pm (CST)
Duration: 30.12s, Total samples = 32.86s (109.11%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 30.07s, 91.51% of 32.86s total
Dropped 258 nodes (cum <= 0.16s)
Showing top 10 nodes out of 71
      flat  flat%   sum%        cum   cum%
    27.77s 84.51% 84.51%     27.82s 84.66%  syscall.Syscall
     0.82s  2.50% 87.01%      0.84s  2.56%  runtime.freedefer
     0.51s  1.55% 88.56%      0.51s  1.55%  runtime.kevent
     0.46s  1.40% 89.96%      0.46s  1.40%  runtime.usleep
     0.16s  0.49% 90.44%      0.53s  1.61%  runtime.gentraceback
     0.15s  0.46% 90.90%      0.21s  0.64%  runtime.pcvalue
     0.08s  0.24% 91.14%     14.25s 43.37%  main.doSomething
     0.04s  0.12% 91.27%     13.88s 42.24%  fmt.Fprintln
     0.04s  0.12% 91.39%      0.88s  2.68%  runtime.deferreturn
     0.04s  0.12% 91.51%      0.23s   0.7%  runtime.scanframeworker
(pprof)
~~~

* 使用 Go 自带的 pprof 工具打开。go tool pprof ~/Downloads/profile。

~~~
# tom @ tom-pc in ~ [22:19:48]
$ go tool pprof ~/Downloads/profile
Main binary filename not available.
Type: cpu
Time: Dec 13, 2018 at 22:16pm (CST)
Duration: 30.12s, Total samples = 32.86s (109.11%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 30.07s, 91.51% of 32.86s total
Dropped 258 nodes (cum <= 0.16s)
Showing top 10 nodes out of 71
      flat  flat%   sum%        cum   cum%
    27.77s 84.51% 84.51%     27.82s 84.66%  syscall.Syscall
     0.82s  2.50% 87.01%      0.84s  2.56%  runtime.freedefer
     0.51s  1.55% 88.56%      0.51s  1.55%  runtime.kevent
     0.46s  1.40% 89.96%      0.46s  1.40%  runtime.usleep
     0.16s  0.49% 90.44%      0.53s  1.61%  runtime.gentraceback
     0.15s  0.46% 90.90%      0.21s  0.64%  runtime.pcvalue
     0.08s  0.24% 91.14%     14.25s 43.37%  main.doSomething
     0.04s  0.12% 91.27%     13.88s 42.24%  fmt.Fprintln
     0.04s  0.12% 91.39%      0.88s  2.68%  runtime.deferreturn
     0.04s  0.12% 91.51%      0.23s   0.7%  runtime.scanframeworker
(pprof)
~~~


* 看到 cpu 占用前 10 的函数，我们可以进行分析优化。
如果看着不直观，可以输入命令 web（需要先安装 graphviz，macOS 下可以 brew install graphviz），输入web,将生成pprofxxx.svg,拖到浏览器查看就可。


- 火焰图 --图形界面
macOS 推荐使用 go-torch 工具。(需要先安装FlameGraph，(
git clone https://github.com/brendangregg/FlameGraph.git,进入cd FlameGraph目录下，执行go-torch 命令)

~~~
# tom @ tom-pc in ~/FlameGraph on git:master o [22:49:56] C:130
$ go-torch -u http://localhost:8081 -t 30
INFO[22:50:02] Run pprof command: go tool pprof -raw -seconds 30 http://localhost:8081/debug/pprof/profile
INFO[22:50:33] Writing svg to torch.svg

# tom @ tom-pc in ~/FlameGraph on git:master x [22:58:45]
$ find ./ -name "torch.svg"
.//torch.svg
~~~
 * 生成 torch.svg 文件，路径：```/Users/tom/FlameGraph/torch.svg```。拖到浏览器打开，

* 纵向y轴代表的是函数调用栈，横向x轴各个方块的宽度代表的是占用cpu时间的比例，越宽代表占用cpu时间越多，方块的颜色是随机的没有实际意义。


###### 2.非web程序--runtime/pprof

* 以查看内存为例，生成heap.profile 文件。使用 go tool pprof 工具进行分析。
~~~
package main

import (
	"log"
    "os"
	"runtime/pprof"
	"time"
    "fmt"
)

func main() {

	heapProfile()
}

func heapProfile() {
	f, err := os.OpenFile("heap.profile", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	time.Sleep(10 * time.Second)

	pprof.WriteHeapProfile(f)
	fmt.Println("end")
}

~~~

* 使用 ``` go tool pprof```

~~~
# tom @ tom-pc in ~/goprojects/src/goroutine-demo [22:34:23] C:1
$ go tool pprof heap.profile
Main binary filename not available.
Type: inuse_space
Time: Dec 13, 2018 at 22:34pm (CST)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 2086.86kB, 100% of 2086.86kB total
      flat  flat%   sum%        cum   cum%
 1184.27kB 56.75% 56.75%  1184.27kB 56.75%  runtime/pprof.StartCPUProfile
  902.59kB 43.25%   100%   902.59kB 43.25%  compress/flate.NewWriter
         0     0%   100%   902.59kB 43.25%  compress/gzip.(*Writer).Write
         0     0%   100%  1184.27kB 56.75%  main.cpuProfile
         0     0%   100%  1184.27kB 56.75%  main.main
         0     0%   100%  1184.27kB 56.75%  runtime.main
         0     0%   100%   902.59kB 43.25%  runtime/pprof.(*profileBuilder).build
         0     0%   100%   902.59kB 43.25%  runtime/pprof.profileWriter
(pprof) 
~~~

###### 总结
* pprof:侧重查看 cpu 及内存占用情况。
* trace:侧重用于跟踪goroutine运行轨迹(延迟、堵塞、系统调用等)。


### <span id="4">四.error处理</span>

* 实现 net.Error的interface。
* 在分布式服务中，根据错误类型，如果是超时，则发起重试，如果是临时则做响应处理，再返回给下游。
~~~go
package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

//包装错误，判断错误类型，如果是因为网络原因导致的超时错误或者临时错误，可以发起增量重试
type ResultErr struct {
	error
	IsTimeOut   bool
	IsTemporary bool
}

func (r *ResultErr) Timeout() bool {

	return r.IsTimeOut
}
func (r *ResultErr) Temporary() bool {

	return r.IsTemporary
}

func main() {

	var isTryCount = 0
Loop:
	err := process()
	if err2, ok := err.(net.Error); ok && err2.Timeout() {
		time.Sleep(1e9)
		isTryCount++
		//重试一次
		if isTryCount > 1 {

		} else {
			goto Loop
		}

	}

	if err != nil {
		fmt.Println("err:")
	}
	fmt.Println("end")
}

func process() (ret net.Error) {

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("i:", i)
			ret = &ResultErr{error: nil, IsTimeOut: true, IsTemporary: false}
			ret.Timeout()
		}(i)
	}
	wg.Wait()

	return ret
}

~~~

### <span id="5">五.何时选择并发<span>
    * IO密集型，例如 读写文件、db、redis 等连接。瓶颈在于IO的等待上，并发只提高CPU的利用效率，反而让很多goroutine处于堵塞状态(GWaiting)。
    * cpu密集型，采用并发，提高cpu计算能力。

### <span id="6">六.map线程不安全<span>
    
### <span id="7">七.组合优于继承<span>
