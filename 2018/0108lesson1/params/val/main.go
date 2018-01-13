//主要演示参数传递
//string 、 int 、 bool 、array 、 slice 、 map 、 chan
package main

import (
	"fmt"
)

func main() {
	fmt.Println("--------------------")

	a := 666
	fmt.Println("a's address is:", &a)
	Foo(a)
	fmt.Println("a:", a)

	fmt.Println("--------------------")

	Foo2(&a)
	fmt.Println("a2:", a)
}

//值传递
func Foo(b int) {
	//只在调用时才分配。调用结束后将释放内存。
	fmt.Println("b's address is :", &b)
	b = 888
	fmt.Println("b:", b)
}

//指针传递
func Foo2(c *int) {
	fmt.Println("c's address is:", c)
	*c = 999
	fmt.Println("c:", *c)
}
