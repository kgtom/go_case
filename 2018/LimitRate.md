
## 介绍两种限速

### 一、现在每秒至多执行三次，超过次数，则丢弃。

~~~go
package main

import (
	"fmt"

	"sync"
	"time"
)

type LimitRate struct {
	currentRate int
	startTime   time.Time
	limitCount  int
	lock        sync.Mutex
}

func (l *LimitRate) IsBeyondLimit() bool {

	ret := true
	l.lock.Lock()
	//先检查是否超过请求次数，如果在1s内超了，则重置，否则返回false
	if l.limitCount == l.currentRate {
		t := time.Now().Sub(l.startTime)
		//fmt.Println("t:", t, time.Second, l.limitCount, l.currentRate)
		if t >= time.Second { //time.Nanosecond
			fmt.Println("newT:", t)
			l.currentRate = 0
			l.startTime = time.Now()
		} else {

			ret = false
		}
	} else {
		l.currentRate++
	}
	defer l.lock.Unlock()
	return ret
}
func (l *LimitRate) SetLimitCount(cnt int) {
	l.startTime = time.Now()
	l.limitCount = cnt
}
func (l *LimitRate) GetLimitCount() int {
	return l.limitCount
}

func main() {
	var wg sync.WaitGroup
	l := new(LimitRate)
	//限流 每秒至多3次请求
	l.SetLimitCount(3)

	for i := 0; i < 5; i++ {
		wg.Add(1)

		go func(i int) {

			if l.IsBeyondLimit() {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}

			wg.Done()
		}(i)
	}
	wg.Wait()
}
~~~

### 二、现在速率，每秒/3的速率执行。

~~~go

//Limit 限速
package main

import (
	"fmt"
	"sync"
	"time"
)

//LimitRate 限速
type LimitRate struct {
	rate       int           //每秒几个
	interval   time.Duration //限速速率
	lastHandle time.Time     //最后一次处理的时间
	lock       sync.Mutex
}

func (l *LimitRate) Limit() bool {
	result := false
	for {
		l.lock.Lock()

		//判断最后一次执行的时间与当前时间的 时间间隔是否大于限速速率
		t := time.Now().Sub(l.lastHandle)
		//大于限速速率，说明执行时间长，重置最后一次提交时间
		if t > l.interval {

			l.lastHandle = time.Now()
			result = true
		}
		l.lock.Unlock()
		if result {
			return result
		}
		//<=限速速率，说明执行太快，需要等待执行
		time.Sleep(l.interval)
	}
}

//SetRate 设置Rate及限速速率
func (l *LimitRate) SetRate(r int) {
	l.rate = r
	//秒/次数=速率333.333ms
	l.interval = time.Microsecond * time.Duration(1000*1000/l.rate)
}

//GetRate 获取Rate
func (l *LimitRate) GetRate() int {
	return l.rate
}

func main() {
	var wg sync.WaitGroup
	l := new(LimitRate)

	//设置每秒/3的速率执行，如果请求太快，则限速，执行time.Sleep(速率)，类似于队列，排队执行。
	l.SetRate(3)

	startTime := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			if l.Limit() {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("总共时间：", time.Since(startTime))
}

~~~

### 3.结合批量执行定时任务

~~~go
for _,v:=range userList {
    
    go func(){
    //使用
        lr.Limit()
        user:=db.user.Update(v)
        
    }()
}


~~~
