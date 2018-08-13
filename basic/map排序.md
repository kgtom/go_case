
### map排序分 按照key 、按照value排序
~~~go

package main

import (
	"fmt"
	"sort"
)

type Node struct {
	Key int
	Val int
}
type NodeList []Node

func (p NodeList) Swap(i, j int) {

	p[i], p[j] = p[j], p[i]

}
func (p NodeList) Len() int { return len(p) }
func (p NodeList) Less(i, j int) bool {
	return p[i].Val < p[j].Val
}
func main() {
	//按照key或者插入 排序
	// a := make(map[int]int)
	// a[1] = 11
	// a[3] = 33
	// a[2] = 22
	// var x []int
	// fmt.Println(x)
	// for k, v := range a {
	// 	fmt.Println(k, v)
	// 	x = append(x, k)
	// }
	// sort.Ints(x)
	// for _, v := range x {
	// 	fmt.Printf("key:%d-->value:%d\n", v, a[v])
	// }

	//按照value排序

	a := make(map[int]int)
	a[1] = 11
	a[2] = 22
	a[4] = 44
	a[3] = 33
	var x NodeList
	for k, v := range a {
		x = append(x, Node{Key: k, Val: v})
	}
	sort.Sort(x)
	for _, v := range x {
		fmt.Printf("key:%d-->value:%d\n", v.Key, v.Val)
	}

}
~~~
