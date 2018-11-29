//面向对象-多态：同一个方法，多用于不同对象，产生不同作用效果。
//举例：父类定义虚函数或者抽象方法，子类重新父类方法，调用的时候根据运行时类型调用各自方法。
//运行时类型相当于golang 类型断言，找出类型，然后调用该类型方法。
~~~go
package main

import (
	"fmt"
)

//类型重命名
type MySlice []interface{}

func NewMyslice() MySlice {

	return make(MySlice, 0)
}

//自定义结构体
type Order struct {
	Id   string
	Name string
}

//定义一个interface
type Comparable interface {
	IsEqual(obj interface{}) bool
}

func CommIsEqual(a, b interface{}) bool {
	if obj, ok := a.(Comparable); ok {
		return obj.IsEqual(b)
	} else {

		return a == b
	}

}
func (order *Order) IsEqual(obj interface{}) bool {

	if ord, ok := obj.(*Order); ok {
		return order.Id == ord.Id
	}
	fmt.Println("unexpected type ")
	return false
}

func (m *MySlice) Add(elem interface{}) {

	for _, v := range *m {

		if CommIsEqual(v, elem) {
			//if v == elem {
			fmt.Println("this elem already exists:", elem)
			return
		}
	}
	*m = append(*m, elem)
	if obj, ok := elem.(*Order); ok {

		fmt.Println("add success:", obj.Name)
	} else {
		fmt.Println("add success:", elem)
	}
}

func (m *MySlice) Remove(elem interface{}) {

	var isFind bool
	for k, v := range *m {
		if CommIsEqual(v, elem) {
			//if v == elem {
			isFind = true
			*m = append((*m)[:k], (*m)[k+1:]...)
		}
	}
	if !isFind {
		fmt.Println("this elem not exists:", elem)
		return
	}
	fmt.Println(*m)

}
func main() {

	//整形切片
	arr := []int{1, 2, 3}
	AddInt(arr, 4)
	AddInt(arr, 1)

	fmt.Println("-------------------------")
	//整形自定义类型切片
	mySlice := NewMyslice()

	mySlice.Add(2)
	mySlice.Add(3)
	mySlice.Add(3)
	mySlice.Remove(2)
	mySlice.Remove(2)
	fmt.Println("-------------------------")

	//字符串切片
	mySlice.Add("hello")
	mySlice.Add("Golang")
	mySlice.Add("hello")
	mySlice.Remove("Golang")
	mySlice.Remove("Golang")

	////自定义类型
	t1 := &Order{Id: "t01", Name: "tom"}
	t2 := &Order{Id: "t02", Name: "jerry"}

	mySlice.Add(t1)
	mySlice.Add(t1)
	mySlice.Add(t2)
	mySlice.Remove(t2)
	mySlice.Remove(t2)

}

func AddInt(arr []int, elem int) []int {

	for _, v := range arr {
		if v == elem {
			fmt.Println("this elem already exists:", elem)
			return arr
		}
		arr = append(arr, elem)

	}
	fmt.Println("add success:", elem)
	return arr
}

~~~
