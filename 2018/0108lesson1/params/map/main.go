package main

import (
	"fmt"
)

func main() {
	m := map[string]string{"中国": "北京"}
	fmt.Printf("m's address is: %p \n", &m)

	m1 := m
	m1["中国"] = "上海"
	fmt.Printf("m1's address is: %p \n", &m1)
	fmt.Println("m:", m)
	fmt.Println("m1:", m)
	//两者都改变，看似是引用传递，实际是值传递，因为地址不同。改变是因为：map内部维护了一个指针，指向真正的存储空间。与slice 一样。
}
