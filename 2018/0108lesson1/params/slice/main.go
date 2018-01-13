package main

import (
	"fmt"
)

func main() {
	fmt.Println("--------值传递--------------")

	s1 := []int{6, 6, 6, 6}
	fmt.Printf("s1's address is: %p \n", &s1)
	Foo(s1)
	fmt.Println("----------------------")
	fmt.Println("s1:", s1)

}
func Foo(s2 []int) {
	fmt.Println("----------------------")
	fmt.Printf("s2_1's address is: %p \n", &s2)
	//s2[0] += 100 //演示指针传递//含了一个指向底层数组的指针
	s2 = []int{8, 8, 8, 8}
	fmt.Println("s2:", s2)
	fmt.Printf("s2_2's address is: %p \n", &s2)

}

//slice 源码:https://golang.org/src/runtime/slice.go?h=slice#L11
