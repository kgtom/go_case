## 题目

给定一个元素都是正整数的数组A ，正整数 L 以及 R (L <= R)。

求连续、非空且其中最大元素满足大于等于L 小于等于R的子数组个数。
~~~
例如 :
输入: 
A = [2, 1, 4, 3]
L = 2
R = 3
输出: 3
解释: 满足条件的子数组: [2], [2, 1], [3].
注意:

L, R  和 A[i] 都是整数，范围在 [0, 10^9]。
数组 A 的长度范围在[1, 50000]。
~~~

[来源](https://leetcode-cn.com/problems/number-of-subarrays-with-bounded-maximum/)

## 代码
~~~go
package main

import (
	"fmt"
)

func main() {

	arr := []int{2, 1, 4, 3}
	ret := numSubarrayBoundedMax(arr, 2, 3)
	fmt.Println("ret:", ret)
}

func numSubarrayBoundedMax(A []int, l, r int) int {

	ret, j, count := 0, 0, 0 // 返回、不合法条件的数组索引、符合条件数组元素个数
	for k, v := range A {
		if v >= l && v <= r {
			count = k - j + 1 //记录当前元素符合<=r的元素个数
			ret += count
		} else if v < l {
			ret += count // <l 当前元素值都可以与count个数里面的元素值组合新子数组满足其最大值在[l,r]之间
		} else {
			j = k + 1 //记录当前不符合条件数组索引
			count = 0 //记录当前符合条件元素个数 为0
		}
	}
	return ret
}

~~~
