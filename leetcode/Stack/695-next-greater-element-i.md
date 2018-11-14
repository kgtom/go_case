## 题目
给定两个没有重复元素的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。找到 nums1 中每个元素在 nums2 中的下一个比其大的值。

nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出-1。
~~~
示例 1:

输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
输出: [-1,3,-1]
解释:
    对于num1中的数字4，你无法在第二个数组中找到下一个更大的数字，因此输出 -1。
    对于num1中的数字1，第二个数组中数字1右边的下一个较大数字是 3。
    对于num1中的数字2，第二个数组中没有下一个更大的数字，因此输出 -1。
示例 2:

输入: nums1 = [2,4], nums2 = [1,2,3,4].
输出: [3,-1]
解释:
    对于num1中的数字2，第二个数组中的下一个较大数字是3。
    对于num1中的数字4，第二个数组中没有下一个更大的数字，因此输出 -1。
注意:

nums1和nums2中所有元素是唯一的。
nums1和nums2 的数组大小都不超过1000。
~~~

[来源](https://leetcode-cn.com/problems/next-greater-element-i/)
## 代码
~~~go
package main

import "fmt"

func main() {

	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}

	ret := nextGreaterElement2(nums1, nums2)
	fmt.Println("ret:", ret)
}

//空间复杂度O(n^2) 双循环，空间复杂度O(1)
func nextGreaterElement(findNums []int, nums []int) []int {

	ret := []int{} // 或者 make([]int,len(findNums))
	for i := 0; i < len(findNums); i++ {

		for j := 0; j < len(nums); j++ {
			if findNums[i] == nums[j] {
				if j+1 <= len(nums)-1 && nums[j+1] > findNums[i] {

					ret = append(ret, nums[j+1])
				} else {
					ret = append(ret, -1)
				}
			}
		}
	}
	return ret
}

//巧用map，时间复杂度O(n) 一个循环就可以
func nextGreaterElement2(findNums []int, nums []int) []int {

	ret := []int{}
	m := make(map[int]int)

	for k, v := range nums {
		m[v] = k
	}
	for _, v1 := range findNums {
		fmt.Println("dd:", m[99])
		j := m[v1] + 1
		if j < len(nums) && v1 < nums[j] {
			ret = append(ret, nums[j])
		} else {
			ret = append(ret, -1)
		}
	}
	return ret
}

~~~
