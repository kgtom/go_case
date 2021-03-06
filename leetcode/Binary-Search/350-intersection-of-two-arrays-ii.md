## 题目
给定两个数组，编写一个函数来计算它们的交集。

示例 1:
~~~
输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2,2]
~~~
示例 2:
~~~
输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [4,9]
~~~
说明：

输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
我们可以不考虑输出结果的顺序。
进阶:

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小很多，哪种方法更优？
如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

## 代码
~~~go

package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}

	ret := intersect(nums1, nums2)
	fmt.Println("ret:", ret)

}

// 方法一：借助map过滤
func intersect(nums1 []int, nums2 []int) []int {

	m := map[int]int{}
	ret := []int{}
	for _, v := range nums1 {
		m[v]++
	}
	for _, v := range nums2 {
		if m[v] > 0 {
			ret = append(ret, v)
		}
	}
	return ret
}

/二分法:前提是 两个数组已排序的。
func intersect2(nums1 []int, nums2 []int) []int {

	ret := []int{}
	for _, v := range nums1 {
		idx := binarySearch(nums2, v)
		if idx >= 0 && idx <= len(nums2) && nums2[idx] == v {

			ret = append(ret, v)
			nums2 = nums2[idx+1:]
		}
	}
	return ret
}
func binarySearch(nums []int, val int) int {

	low, high := 0, len(nums)-1
	for low < high {
		mid := (low + high) / 2
		if nums[mid] == val {
			return mid
		} else if nums[mid] > val {
			high = mid
		} else {
			low = mid + 1
		}

	}
	return high
}
~~~
