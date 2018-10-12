## 题目

给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:
~~~
输入: [0,1,0,3,12]
输出: [1,3,12,0,0]

~~~
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。

[来源](https://leetcode-cn.com/problems/move-zeroes/)

## 代码
~~~go
package main

import (
	"fmt"
)

func main() {

	arr := []int{0, 1, 3, 0}
	moveZeroes(arr)
	fmt.Println("ret:", arr)

}
func moveZeroes(nums []int) {

	//将非零数，使用j重塑数组，然后将nums[j:]重置为0
	i, j := 0, 0
	for i < len(nums) {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++ //有效非零索引值
		}

		i++

	}
	//重置0
	for j < len(nums) {
		nums[j] = 0
		j++
	}

}

~~~
