## 题目
峰值元素是指其值大于左右相邻值的元素。

给定一个输入数组 nums，其中 nums[i] ≠ nums[i+1]，找到峰值元素并返回其索引。

数组可能包含多个峰值，在这种情况下，返回任何一个峰值所在位置即可。

你可以假设 nums[-1] = nums[n] = -∞。
~~~

示例 1:

输入: nums = [1,2,3,1]
输出: 2
解释: 3 是峰值元素，你的函数应该返回其索引 2。
示例 2:

输入: nums = [1,2,1,3,5,6,4]
输出: 1 或 5 
解释: 你的函数可以返回索引 1，其峰值元素为 2；
     或者返回索引 5， 其峰值元素为 6。
     
 ~~~
 
说明:

你的解法应该是 O(logN) 时间复杂度的。

[来源](https://leetcode-cn.com/problems/find-peak-element/)

## 代码
~~~go
package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 2, 1, 3, 5, 6, 4}
	ret := findPeakElement3(arr)
	fmt.Println("ret:", ret)
}

//常规方法，遍历比较 找到第一个峰值就return
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

//二分法查找--迭代版本，只能查找到最后一个，不能找到第一个峰值。
//二分查找要比较的是 target 元素，本题的 target 元素是 mid+1元素，即 nums[mid] 与 nums[mid+1] 的比较
//mid>mid+1说明峰值在左边，反之峰值在右边
func findPeakElement2(nums []int) int {
	l := 0
	r := len(nums) - 1
	mid := (l + r) / 2

	for l < r {
		fmt.Println("mid:", mid, "l:", l, "r:", r)
		if nums[mid] < nums[mid+1] {

			l = mid + 1
		} else {
			r = mid
		}
		mid = (l + r) / 2
	}

	return l
}

//二分法查找--递归版本
func findPeakElement3(nums []int) int {
	return search(nums, 0, len(nums)-1)
}

func search(nums []int, l, r int) int {
	mid := (l + r) / 2

	for l < r {

		//mid>mid+1则找到了
		if nums[mid] > nums[mid+1] {
			return search(nums, l, mid)
		}
		//mid<=mid+1
		return search(nums, mid+1, r)

	}
	return l
}

//循环遍历一遍O(n), 只判断nums[i]>nums[i+1]就是第一个极值,如果没有则说明是递增序列，返回最后索引
func findPeakElement4(nums []int) int {

	for i := 0; i < len(nums)-1; i++ {
		//
		if nums[i] > nums[i+1] {
			return i
		} else {
			//nums[i]<=nums[i+1] 说明前面都是一个递增序列，只要有一个 nums[i]>nums[i+1]的就是峰值
		}
	}
	return len(nums) - 1
}

~~~
