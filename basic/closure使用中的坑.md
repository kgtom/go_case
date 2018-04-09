

### closure 的介绍
***
```go
package main

import (
	"fmt"
)

func main() {

	var i int

	//closure1

	f := func() int {

		i++
		x := i
		fmt.Println("x:", x) //1
		return x

	}

	fmt.Println("i:", i) //0 未调用f()

	//fmt.Println("x2:", x) //closure内变量，访问不到。

	f()

	//closure2

	{

		z := i
		fmt.Println("z:", z) //1

	}

	//fmt.Println("z2:", z)//closure内变量，访问不到。

	fmt.Println("i2:", i) //1 调用f()后i++

	fmt.Println("end")

}


```

### closure 坑
***
```go
  func main() {

	a := []string{"tom", "lili", "lucy"}

	for _, v := range a {
		go func() {
			fmt.Println("v:", v)
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("end")

}

```
`输出 三个 lucy,原因：执行循环体，获取最后一次循环的值`

改进方案一：
***
```go
func main() {

	a := []string{"tom", "lili", "lucy"}

	for _, v := range a {
		vv := v//拷贝一份新的
		go func() {
			fmt.Println("v:", vv)
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println("end")

}
```
改进方案二：
***
```go
func main() {

	a := []string{"tom", "lili", "lucy"}

	for _, v := range a {

		go func(vv string) {
			fmt.Println("v:", vv)
		}(v)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("end")

}
```

