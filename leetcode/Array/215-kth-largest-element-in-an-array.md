
## 题目


[来源](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/)
## 代码

~~~go

package main

import (
	"fmt"
)

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	ret := findKthLargest(nums, 2)
	fmt.Println("ret:", ret)
}
func findKthLargest(nums []int, k int) int {
	return innerFind(nums, k, 0, len(nums)-1)
}

//快排，不一定完成所有排序，只需要找到中轴点索引与 k的大小。
// 如果正好pivotIdx==k,则nums[pivotIdx]，如果 pivotIdx>k，则需要从大的那部分继续快排，找到pivotIdx==k的数
//同理如果pivotIdx<k。则需要从小的部分继续快排。
func innerFind(nums []int, k int, start, end int) int {

	for {
		pivotIdx := partition(nums, start, end)
		//fmt.Println(pivotIdx)
		if pivotIdx == k {
			return nums[pivotIdx]
		} else if pivotIdx < k {
			start = pivotIdx + 1
		} else {
			end = pivotIdx - 1
		}
	}
}

//二分法，左边小于中轴值，右边大于中轴值，返回中轴索引
func partition(nums []int, left, right int) int {
	pivotVal := nums[left]
	pivotIdx := left
	for left < right {
		for left < right && nums[left] <= pivotVal {
			left++
		}
		for left < right && nums[right] > pivotVal {
			right--
		}
		if left == right {
			break
		}

		nums[left], nums[right] = nums[right], nums[left]

	}
	nums[left], nums[pivotIdx] = nums[pivotIdx], nums[left]
	return left
}

~~~
