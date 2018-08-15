## 题目
给定排序数组和目标值，如果找到目标，则返回索引。如果没有，请返回索引按顺序插入的索引。

您可以假设数组中没有重复项。

例1：

~~~
 输入： [1,3,5,6]，5
 输出： 2
~~~
 
例2：

~~~
 输入： [1,3,5,6]，2
 输出： 1
~~~
 
例3：
~~~
 输入： [1,3,5,6]，7
 输出： 4
 ~~~
 
例4：

~~~

 输入： [1,3,5,6]，0
 输出： 0
~~~

[来源](https://leetcode.com/problems/search-insert-position/description/)

## 代码

~~~go
package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 3, 4, 5, 9}
	r := searchInsert(arr, 8)
	fmt.Println("r:", r)
}

//二分查找方法实现
func searchInsert(arr []int, target int) int {

	low := 0
	high := len(arr) - 1
	for low <= high {

		mid := (low + high) >> 1 //(low + high) / 2 mid溢出
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}


~~~
