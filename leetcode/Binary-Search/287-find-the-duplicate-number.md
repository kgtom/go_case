## 题目

给定一个包含 n + 1 个整数的数组 nums，其数字都在 1 到 n 之间（包括 1 和 n），可知至少存在一个重复的整数。假设只有一个重复的整数，找出这个重复的数。

示例 1:
~~~
输入: [1,3,4,2,2]
输出: 2
~~~

示例 2:
~~~
输入: [3,1,3,4,2]
输出: 3
~~~

说明：

不能更改原数组（假设数组是只读的）。
只能使用额外的 O(1) 的空间。
时间复杂度小于 O(n2) 。
数组中只有一个重复的数字，但它可能不止重复出现一次。
## 代码
~~~
package main

import "fmt"

func main() {
	nums := []int{3, 1, 3, 4, 2}
	ret := findDuplicate(nums)
	fmt.Println("ret:", ret)

}

//利用二分法,time:O(nlogn)，与10个苹果9个抽屉放置，找到一个抽屉放2个苹果，思路一样。折半找。
// 如果小于等于mid的数的个数cnt> mid，那么重复的数字一定出现在[l，mid]之间，反之在[mid,higt]中。
func findDuplicate(nums []int) int {
	low, high := 0, len(nums)-1

	for low < high {
		mid := (high + low) / 2
		cnt := 0 //记录个数
		for i := 0; i < len(nums)-1; i++ {

			if nums[i] <= mid {
				cnt++
			}

		}
		//说明在前半部分
		if cnt > mid {
			high = mid - 1
		} else if cnt == mid {
			return mid
		} else {
			//说明在后半部分
			low = mid + 1
		}
	}

	return low
}
~~~
