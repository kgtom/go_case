## 题目
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。
~~~
示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
~~~


## 代码
~~~go
package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 2, 3, 1}
	ret := findPeakElement(arr)
	fmt.Println("ret:", ret)
}

//遍历比较
func findPeakElement(nums []int) int {

	for i := 0; i < len(nums); i++ {
		//注意索引边界问题
		if i > 0 && i < len(nums)-1 && nums[i] > nums[i-1] && nums[i+1] < nums[i] {
			return i
		}
		if i > 0 && i == len(nums)-1 && nums[i] > nums[i-1] {
			return i
		}
	}
	return -1
}

~~~
