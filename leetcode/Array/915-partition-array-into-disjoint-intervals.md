## 题目

给定一个数组 A，将其划分为两个不相交（没有公共元素）的连续子数组 left 和 right， 使得：

left 中的每个元素都小于或等于 right 中的每个元素。
left 和 right 都是非空的。
left 要尽可能小。
在完成这样的分组后返回 left 的长度。可以保证存在这样的划分方法。

 
~~~
示例 1：

输入：[5,0,3,8,6]
输出：3
解释：left = [5,0,3]，right = [8,6]
~~~

~~~
示例 2：

输入：[1,1,1,0,6,12]
输出：4
解释：left = [1,1,1,0]，right = [6,12]
~~~

提示：

2 <= A.length <= 30000
0 <= A[i] <= 10^6
可以保证至少有一种方法能够按题目所描述的那样对 A 进行划分。
 
[来源](https://leetcode-cn.com/problems/partition-array-into-disjoint-intervals/description/)

## 代码
~~~go
package main

import "fmt"

func main() {

	arr := []int{3, 1, 1, 0, 6, 12}
	ret := partitionDisjoint(arr)
	fmt.Println("ret", ret)
}

//时间复杂度是O(N)，空间复杂度是O(1)
func partitionDisjoint(nums []int) int {

	disjoint := 0
	max := nums[0]

	for i := 1; i < len(nums); i++ {

		if nums[i] < max {

			disjoint = i
		}
	}

	//第一个是最大的数
	if disjoint == len(nums)-1 {

		fmt.Println("left:", nums[1:])
		fmt.Println("right:", nums[0])
		return disjoint
	}
	fmt.Println("left:", nums[:disjoint+1])
	fmt.Println("right:", nums[disjoint+1:])
	return disjoint + 1
}

~~~
