## 题目
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。



在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 5
输出:
[
     [1],
    [1,1],
   [1,2,1],
  [1,3,3,1],
 [1,4,6,4,1]
]
[来源](https://leetcode-cn.com/problems/pascals-triangle/description/)
## 代码
~~~go
package main

import (
	"fmt"
)

func main() {

	ret := generate(4)
	fmt.Println("ret:", ret)

}
func generate(numRows int) [][]int {
	ret := make([][]int, 0, 0) //或者 [][]int{}
	if numRows == 0 {
		return ret
	}

	ret = append(ret, []int{1}) //第一组数据初始化
	if numRows == 1 {
		return ret
	}

	for i := 1; i < numRows; i++ {
		next := genNext(ret[i-1])

		ret = append(ret, next) //拼接所有数据
	}

	return ret
}

func genNext(prev []int) []int {
	next := make([]int, 1, len(prev)+1)

	next = append(next, prev...) //将上一组数据拿过来，用于计算下一组数据
	fmt.Println("len(next):", len(next), "prev:", prev, "next:", next)
	for i := 0; i < len(next)-1; i++ {
		next[i] += next[i+1] //相邻相加，改变值

	}
	fmt.Println("next:", next) //返回组合完成的数据
	return next
}

~~~
