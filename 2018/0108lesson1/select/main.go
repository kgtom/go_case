package main

import (
	"fmt"
	"time"
)

func main() {

	tick := time.Tick(500 * time.Millisecond)

	tick2 := time.After(1000 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick....")
		case <-tick2:
			fmt.Println("tick2...")
			//执行完成后返回，不然程序一直跑下去
			return
		default:
			//默认
			fmt.Println("dafault...")
			//没有Sleep，程序会优先匹配default
			time.Sleep(250 * time.Millisecond)
		}
	}

}
