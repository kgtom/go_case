
### 第一  从close的unbuffered channel上执行读操作，返回channel对应类型的零值，但写入会panic。读写都不会堵塞。

```go
ch  :=  make(chan  bool)
close(ch)
v  :=  <-ch
fmt.Println("ch:", v) //false

v, ok  :=  <-ch
fmt.Println("v:", v, "ok:", ok) //false,false

ci  :=  make(chan  int)
close(ci)

fmt.Println("ch int:", <-ci) //0 channel对应类型的零值
ci <-  2  //send on closed channel
```

### 第二 从close带buffered channel,可以读取，返回对应类型的零值，但写入会panic。读写都不会堵塞。

```go
ch  :=  make(chan  int, 2)
ch <-  1
ch <-  2
close(ch)
fmt.Println("ch:", <-ch) //1
fmt.Println("ch:", <-ch) //2
v, ok  :=  <-ch
fmt.Println("v:", v, "ok:", ok) //0,false channel对应类型的零值
ch <-  2  //send on closed channel
```

### 第三 没有初始化channel(nil)读写都会堵塞

```go
var  ch  chan  int
ch <-  1  //fatal error: all goroutines are asleep 
– deadlock!
```
```go
var  ch  chan  int
<-ch //fatal error: all goroutines are asleep 
– deadlock!
```
### 第四 nil channel妙用
先看一个死循环的例子：本想依次输出 1 ，3 数字。

```go 
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(time.Second * 1)
		ch1 <- 1
		close(ch1)
	}()

	go func() {

		time.Sleep(time.Second * 2)
		ch2 <- 3
		close(ch2)
	}()

	for {

		select {
		case v := <-ch1:
			fmt.Println("ch1:", v)
		case v := <-ch2:
			fmt.Println("ch2:", v)
		}
	}
}


```
`输出结果： 1，0，0，0(无数0) 原因：
1.close(ch1)后获取到都是 int类型的零值 0；
2.case顺序执行 `

改正方案：
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	ch2 := make(chan int)

	go func() {

		time.Sleep(time.Second * 1)

		ch1 <- 1

		close(ch1)

	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch2 <- 3
		close(ch2)
	}()

	for {

		select {
		case v, ok := <-ch1:
			if ok {
				fmt.Println("ch1:", v)
			} else {
				ch1 = nil
			}
		case v, ok := <-ch2:
			if ok {
				fmt.Println("ch2:", v)
			} else {
				ch2 = nil
			}
		}

		if ch1 == nil && ch2 == nil {
			break
		}
	}
}


```
`ch1=nil ,堵塞在case这个分支，然后ch2才有机会运行。`
