## 题目
给定长度为 2n 的数组, 你的任务是将这些数分成 n 对, 例如 (a1, b1), (a2, b2), ..., (an, bn) ，使得从1 到 n 的 min(ai, bi) 总和最大。

示例 1:
~~~
输入: [1,4,3,2]

输出: 4
解释: n 等于 2, 最大总和为 4 = min(1, 2) + min(3, 4).
~~~

提示:

n 是正整数,范围在 [1, 10000].
数组中的元素范围在 [-10000, 10000].


[来源:](https://leetcode-cn.com/problems/array-partition-i/description/)

## 代码


~~~go
package main

import (
	"fmt"
)

//排序后，相邻一组，然后取出每一组第一个数字相加即可。
func main() {

	arr := []int{1, 9, 4, 2}
	ret := arrayPairSum(arr)
	fmt.Println("ret:", ret)

}
func arrayPairSum(nums []int) int {
	var newArr = [10]int{} //最大不超9的元素值
	for k, v := range nums {
		fmt.Println("v:", v, "k:", k)
		newArr[v]++
	}
	fmt.Println(newArr)

	ret := 0
	isSum := true
	for k, v := range newArr {

		for v > 0 {
			if isSum {
				ret += k
			}
			isSum = !isSum //间隔相加
			v--
		}
	}

	return ret
}

~~~


