## 题目

给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。
~~~
示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
~~~

[来源](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/description/)



## 代码
~~~go
package main

import "fmt"

func main() {

	nums := []int{5, 7, 7, 8, 9, 8, 10}

	ret := searchRange2(nums, 8)
	fmt.Println("ret", ret)
}

//二分法--循环迭代
func searchRange(nums []int, target int) []int {
	left, right, mid := 0, len(nums)-1, 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			if nums[left] < target {
				left++
			}
			if nums[right] > target {
				right--
			}
			if nums[left] == nums[right] {
				return []int{left, right}
			}
		}
	}
	return []int{-1, -1}
}

//二分法，分为两端
func searchRange2(nums []int, target int) []int {
	// 查看target是否存在与nums中
	idx := BinarySearch(nums, target)
	if idx == -1 {
		return []int{-1, -1}
	}

	// 利用二分法，查找第左边索引
	left := idx
	for {
		l := BinarySearch(nums[:left], target)
		if l == -1 {
			break
		}
		left = l
	}

	// 利用二分法，查找右边索引
	right := idx
	for {
		r := BinarySearch(nums[right+1:], target)
		if r == -1 {
			break
		}

		right += r + 1
	}

	return []int{left, right}
}

// 二分查找法
func BinarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	var mid int

	if l <= r {
		//不使用(l+r)/2，避免l 与 r 最大值时，越界问题
		mid = l + (r-l)/2
		if nums[mid] > target {
			r = mid - 1
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			return mid
		}

	}
	return -1
}

~~~



