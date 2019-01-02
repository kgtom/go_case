
## 题目
给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。


示例 1:
~~~
输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
~~~
示例 2:
~~~
输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1
~~~

[来源](https://leetcode-cn.com/problems/binary-search/)


## 代码

~~~go
package main

import "fmt"

func main() {
	arr := []int{-1, 0, 3, 5, 9, 12}
	ret := search(arr, 9)
	fmt.Println("ret:", ret)

}

//迭代版
func search(nums []int, target int) int {

	low := 0
	hight := len(nums) - 1
	for low < hight {

		mid := (low + hight) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			hight = mid - 1
		}
	}

	return -1
}


~~~
