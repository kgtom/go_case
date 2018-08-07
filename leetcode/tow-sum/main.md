


## 题目

`https://leetcode-cn.com/problems/two-sum/description/
给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。

你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
`

##  代码
~~~ go
/*
 无缓冲channel，因为无缓冲，同步作用。但会面对未知的超时的情况。读写都会堵塞，当数据未发送(未写)的时候，读堵塞；当发送(写)入时，未被读取，写会堵塞。
     特别注意：当channel中有数据，未被消费时，接收发生在发送之前。
 有缓冲channel,缓冲满时，写堵塞，缓冲空的时候，读堵塞。
  goroutine泄露：开启goroutine 堵塞，无法退出。
*/
package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 3, 4, 3, 5}
	twoNums(arr, 6)

}

func twoNums(nums []int, targetValue int) []int {
	m := map[int]int{}
	idxArr := []int{}
	//转化成map
	for k, v := range nums {
		m[v] = k
	}
	fmt.Println(m)

	//从map 中查找
	for k, v := range m {
		left := targetValue - k

		if k1, ok := m[left]; ok {
			//fmt.Println("idx:", k1)

			idxArr = []int{v, k1}

			break
		}

	}
	fmt.Println("idxArr:", idxArr)

	return idxArr
}


~~~
