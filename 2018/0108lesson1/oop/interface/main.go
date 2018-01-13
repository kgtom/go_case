//接口的多种不同的实现方式（表现不同的行为）即为多态本质，一种策略模式的体现
package main

import (
	"fmt"
)

type areaer interface {
	area() int //面向接口原则
}

type rect struct {
	width, height int
}

type square struct {
	side int
}

func (r rect) area() int {
	return r.height * r.width
}

func (s square) area() int {
	return s.side * s.side
}

func main() {
	var i areaer

	i = rect{4, 3}
	fmt.Println(i.area())

	i = square{6}

	fmt.Println(i.area())
}
