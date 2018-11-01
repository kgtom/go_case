## 题目
给定一个整数数组 a，其中1 ≤ a[i] ≤ n （n为数组长度）, 其中有些元素出现两次而其他元素出现一次。

找到所有出现两次的元素。

你可以不用到任何额外空间并在O(n)时间复杂度内解决这个问题吗？

示例：
~~~
输入:
[4,3,2,7,8,2,3,1]

输出:
[2,3]
~~~

[来源](https://leetcode-cn.com/problems/find-all-duplicates-in-an-array/)
## 代码
~~~go

package main

import "fmt"

func main() {

	arr := []int{4, 3, 2, 7, 8, 2, 3, 1}
	ret := findDuplicates(arr)
	fmt.Println("ret", ret)
}

//不用到任何额外空间并在O(n)时间复杂度内解决
func findDuplicates(nums []int) []int {

	mapKey := make(map[int]int)
	i := 0

	for _, v := range nums {
		mapKey[v]++

		//记录重复数，重置nums数据
		if mapKey[v] > 1 {

			nums[i] = v
			i++
		}

	}
	//fmt.Println("nums:", nums[:i])
	return nums[:i]

}

//使用额外空间并在O(n)时间复杂度内解决
func findDuplicates2(nums []int) []int {

	mapKey := make(map[int]int)
	dup := []int{}

	for _, v := range nums {
		mapKey[v]++

		//过滤
		if mapKey[v] > 1 {
			dup = append(dup, v)
		}
	}
	return dup

}

~~~
