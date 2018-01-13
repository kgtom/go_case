package main

import "fmt"

//定义Rect结构体
type Rect struct {
	x, y          float64
	width, height float64
}

//为Rect结构体增加计算面积的方法
func (r *Rect) Area() float64 {
	return r.width * r.height
}

//创建初始化Rect的函数---模拟“构造函数”
func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

//声明Base
type Base struct {
	Name string //访问权限：大写字母开头
}

//为Base添加funca和funcb
func (b *Base) funca() {
	fmt.Println("funca:", b.Name)
}
func (b *Base) funcb() {
	fmt.Println("funcb", b.Name)
}

//继承Base的Children
type Children struct {
	Base      //匿名字段就是用来实现继承的，golang成为：匿名组合
	ClildName string
}

//重写 funcb
func (c *Children) funcb() {
	fmt.Println("funcb:", c.Name, c.ClildName)
}

//为children添加funcc方法
func (c *Children) funcc() {
	fmt.Println("funcc", c.Name, c.ClildName)
}
func main() {

	// //测试new出来的结构体是什么类型
	// testRect1 := new(Rect)
	// testRect1.x = 500
	// testRect2 := testRect1
	// testRect2.x = 300
	// fmt.Println(testRect1, testRect2) //&{300 0 0 0} &{300 0 0 0}

	// fmt.Println("-----------rect struct--------------")
	// //创建初始化Rect类型
	// rect1 := new(Rect)
	// fmt.Println(rect1)

	// rect2 := &Rect{}
	// fmt.Println(rect2)

	// rect3 := &Rect{0, 0, 100, 200}
	// fmt.Println(rect3)

	// rect4 := &Rect{width: 200, height: 400}
	// fmt.Println(rect4)

	// rect5 := NewRect(20, 20, 200, 200)
	// fmt.Println(rect5)
	// fmt.Println("-----------base--------------")
	//Base
	b := &Base{"我是Base"}
	b.funca()
	b.funcb()
	fmt.Println("-----------child--------------")
	c := &Children{*b, "我是child"}
	c.funca() //自己没有，使用父类
	c.funcb() //重写基类,只用自己
	c.funcc() //子类独有，体现里氏替换原则（demo:oop/child）
}
