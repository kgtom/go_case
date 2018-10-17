## 题目

[来源](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/)
## 代码

~~~go
package main

import (
	"fmt"
)

func main() {

	nums := []int{1, 1, 2}
	ret := removeDuplicates(nums)
	fmt.Println(ret)
}

//空间复杂度 O(1),时间复杂度O(n)
func removeDuplicates(nums []int) int {
	currIdx := 0
	for i := 0; i < len(nums)-1; i++ {

		//跳过重复元素，只记录有效元素
		if nums[i] != nums[i+1] {
			nums[currIdx] = nums[i]
			currIdx++
		}

	}
	fmt.Println("nums:", nums[currIdx:]) //有效数组
	return currIdx + 1

}


~~~
