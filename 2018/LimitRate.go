
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
