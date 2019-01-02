
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
	ret := search(arr, 9) //search2
	fmt.Println("ret:", ret)

}

//迭代版
func search(nums []int, target int) int {

	low := 0
	high := len(nums) - 1
	for low < high {

		mid := (low + high) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}


//递归版
func search2(nums []int, target int) int {

	return innerSearch(nums, target, 0, len(nums)-1)

}
func innerSearch(nums []int, target, low, high int) int {
	mid := (low + high) / 2
	for low <= high {

		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			return innerSearch(nums, target, low, mid-1)
		} else {
			return innerSearch(nums, target, mid+1, high)
		}
	}
	return -1
}

~~~
