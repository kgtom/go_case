## 题目

给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。

注意你不能在买入股票前卖出股票。
~~~
示例 1:

输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
示例 2:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
~~~

[来源](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)


## 代码
~~~go
package main

import (
	"fmt"
)

func main() {

	arr := []int{7, 1, 5, 3, 6, 4}
	ret := maxProfit(arr)
	fmt.Println("ret:", ret)
}

//动态规划，从前往后，找到最小值与当前遍历值之间的差值，差值越大，利润越大
//时间复杂度 O(n) 空间复杂度O(1)
func maxProfit(prices []int) int {

	ret := 0
	minPrice := prices[0] //默认prices[0]最小

	for i := 1; i < len(prices); i++ {

		//找到最小
		minPrice = min(minPrice, prices[i])

		//找到当前元素与最小的差值，其中最大的
		ret = max(ret, prices[i]-minPrice)
		//fmt.Println(minPrice, ret)
	}

	return ret
}

func min(a, b int) int {

	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

~~~
