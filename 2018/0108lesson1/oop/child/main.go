package main

import (
	"fmt"
)

type A struct {
}

type B struct {
	A //B is-a A
}

func save(a *A) {
	//do something
	fmt.Println("save")
}

func main() {
	b := B{}
	save(&b) //! b IS NOT A 不符合里氏替换原则
}
