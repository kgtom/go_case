## 题目

给定一个非负整数数组 A，返回一个由 A 的所有偶数元素组成的数组，后面跟 A 的所有奇数元素。

你可以返回满足此条件的任何数组作为答案。

 

示例：
~~~
输入：[3,1,2,4]
输出：[2,4,3,1]
输出 [4,2,3,1]，[2,4,1,3] 和 [4,2,1,3] 也会被接受。
~~~

[来源](https://leetcode-cn.com/problems/sort-array-by-parity/)

## 代码
~~~go
package main

import (
	"fmt"
)

func main() {

	ret := sortArrayByParity([]int{1, 3, 5, 4, 7, 8})
	fmt.Println("ret:", ret)

	//ret: [1 3 5 4 7 8]
}

func sortArrayByParity(nums []int) []int {

	i, j := 0, len(nums)-1
	// 采用 收尾不断内收，两面夹击
	for {
		//偶数
		for i < j && nums[i]%2 == 0 {
			i++
		}
		//奇数
		for i < j && nums[j]%2 == 1 {
			j--
		}

		if i > j {
			nums[i], nums[j] = nums[j], nums[i] //奇偶临界点交换数据，其它位置按兵不动
		} else {
			break
		}

	}
	return nums
}

~~~
