
~~~ go

package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {

		wg.Add(1)
		//调用cal()之前add,如果放在cal()后面，将"panic: sync: negative WaitGroup counter"
		cal(i, &wg)
	}

	wg.Wait()

}

func cal(i int, wg *sync.WaitGroup) {
	fmt.Println(i * 2)
	wg.Done()

}

~~~
