### 1.CGO 价值：
cgo 可以继承c\c++近半个实际的软件积累
### 2.CGO场景：
通过OpenGL或者OpenCL 使用显卡的计算能力
通过OpenCV来进行图像分析
通过Go 编写Python 扩展
通过Go编写移动应用
### 3.快速入门
* hello cgo
~~~go
package main

import "C" //表示启用CGo
func main(){
println("hello cgo")
}

~~~
需要安装go语言编译器和cgg编译器后，执行 go run main.go

ps：
面向纯c接口的Go语言编程。
GoString 也是一种c字符串
Go的一切都可以用c去理解。

### 4.类型转换
编程：数据结构 （类型） +算法（函数）

go字符串和切片底层：指针+长度、指针+长度+容量

字符串 与切换时兼容的，实际指针+长度
go:数值类型 int32-->uintptr-->unsafe.Pointer-->其它类型 *C.char ,反之亦然。
go:结构体转换 p-->unsafe.Pointer--->q
go:[]x 和[] y 转换：所有切片拥有相同的头部 reflect.SliceHeader，重构切片头部。

### 5.函数调用
* go-->c：通过improt "c"

* c-->go:通过export 导出go函数

~~~go
import "c"
//export int GoAdd(int a,int b);
//#include "add.h"
func main(){

c.GoAdd(1,1) //是go 导出函数
c.c_add(1,2) //是c定义的函数，在add.h头文件引用
}
~~~
* c<--->go：相互调用，例如：Go main-->C:go_add()--->Go：GoAdd
~~~
func main(){
c.c_add(1,2)
}
//export GoAdd
func GoAdd(a,b  c.int) c.int {
return a+b
}
~~~

### 6.CGo内部机制
慢，内部调用流程多，各种内存数据转换
### 7.内存模型--->理论武装代码


>reference:
[gopherChina2018](https://www.youtube.com/watch?v=FwyEP9XeOhY)





