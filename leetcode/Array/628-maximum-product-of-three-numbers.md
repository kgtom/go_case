## 题目
给定一个整型数组，在数组中找出由三个数组成的最大乘积，并输出这个乘积。

示例 1:

输入: [1,2,3]
输出: 6
示例 2:

输入: [1,2,3,4]
输出: 24
注意:

给定的整型数组长度范围是[3,104]，数组中所有的元素范围是[-1000, 1000]。
输入的数组中任意三个数的乘积不会超出32位有符号整数的范围。

[来源](https://leetcode-cn.com/problems/maximum-product-of-three-numbers/)

## 代码


~~~go
package main

import (
	"fmt"
	"math"
)

func main() {

	arr := []int{1, 2, 3, 4}
	ret := maximumProduct(arr)
	fmt.Println("ret:", ret)
}

//时间复杂度为O(n)
//最主要考虑正负数，要是全部为正数，则最大前三位，如果有负数，则比较后两个与最大正数积

func maximumProduct(a []int) int {
	//巧用最小值初始化最大值，存放最大值
	max := []int{math.MinInt32, math.MinInt32, math.MinInt32}

	//巧用最大值初始化存放最小值
	min := []int{math.MaxInt32, math.MaxInt32}

	for i := 0; i < len(a); i++ {
		//比较，逐个迭代
		if a[i] > max[2] {
			max[0], max[1], max[2] = max[1], max[2], a[i]
		} else if a[i] > max[1] {
			max[0], max[1] = max[1], a[i]
		} else if a[i] > max[0] {
			max[0] = a[i]
		}
		//查找最小的
		if a[i] < min[1] {
			min[0], min[1] = min[1], a[i]
		} else if a[i] < min[0] {
			min[0] = a[i]
		}
	}

	//三个正数最大值
	ret := max[0] * max[1] * max[2]

	//两个最小*一个最大正数(考虑负数的情况)
	if ret < max[2]*min[0]*min[1] {
		return max[2] * min[0] * min[1]
	}

	return ret
}

~~~
