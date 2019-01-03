
## 题目
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:
~~~
输入: 4
输出: 2
~~~
示例 2:
~~~
输入: 8
输出: 2
说明: 8 的平方根是 2.82842..., 
     由于返回类型是整数，小数部分将被舍去。
 ~~~
     
## 代码
~~~go
package main

import "fmt"

func main() {
	ret := mySqrt(8)
	fmt.Println("ret:", ret)

}

//二分查找思路，迭代区间值的判断
func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	low := 1
	high := x

	for low < high {
		mid := (low + high) / 2
		//区间的判断，小于等于且不大于+1的值
		if mid*mid <= x && x < (mid+1)*(mid+1) {
			return mid
		} else if mid*mid < x {
			low = mid
		} else {
			high = mid
		}
	}
	return low
}
~~~
