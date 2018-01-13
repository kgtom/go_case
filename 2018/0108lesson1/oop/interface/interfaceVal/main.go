/*Go 语言规范定义了接口方法集的调用规则：
类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
类型 T 的可调用方法集包含接受者为 T 的所有方法,不包含 *T
*/
package main

import (
	"fmt"
)

type Integer int

func (a Integer) Less(b Integer) bool {

	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
	fmt.Println("Add:", *a)

}

type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}

func main() {
	var a Integer = 1
	var b Integer = 2
	var c LessAdder = &a
	fmt.Println("c.Less():", c.Less(b))
	c.Add(b)
	//var b1 LessAdder = a //错误地方
}
