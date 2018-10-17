
## 题目

给定一个数组 nums 和一个值 val，你需要原地移除所有数值等于 val 的元素，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

示例 1:

给定 nums = [3,2,2,3], val = 3,

函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。

你不需要考虑数组中超出新长度后面的元素。
[来源](https://leetcode-cn.com/problems/remove-element/)

## 代码

~~~go
package main

import (
	"fmt"
)

func main() {

	nums := []int{3, 2, 2, 3}
	ret := removeElement2(nums, 3)
	fmt.Println(ret)
}

//空间复杂度 O(1),时间复杂度O(n)
func removeElement(nums []int, val int) int {
	currIdx := 0 //记录不同数出现的个数
	for i := 0; i < len(nums); i++ {

		if nums[i] != val {
			nums[currIdx] = nums[i]
			currIdx++
		}

	}
	fmt.Println("nums:", nums[:currIdx]) //有效数组
	return currIdx

}

//空间复杂度 O(1),时间复杂度O(n) 适合 数组删除元素少，此方法效率更高。
func removeElement2(nums []int, val int) int {
	currIdx := 0
	n := len(nums)
	for currIdx < n {

		if nums[currIdx] == val {
			//当前的与最后一个交换，让总长度减1
			nums[currIdx] = nums[n-1]
			n--
		} else {
			currIdx++
		}
	}
	fmt.Println("nums:", nums[:n])
	return n
}


~~~
