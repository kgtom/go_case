go run pprof.go ,打开 http://localhost:8080/debug/pprof/

package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

var quit chan struct{} = make(chan struct{})

func f() {

	<-quit
}

func main() {

	for i := 0; i < 10000; i++ {
		fmt.Println(i)
		go f()
	}

	http.ListenAndServe(":8080", nil)
}
