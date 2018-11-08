
## 题目

数组的每个索引做为一个阶梯，第 i个阶梯对应着一个非负数的体力花费值 cost[i](索引从0开始)。

每当你爬上一个阶梯你都要花费对应的体力花费值，然后你可以选择继续爬一个阶梯或者爬两个阶梯。

您需要找到达到楼层顶部的最低花费。在开始时，你可以选择从索引为 0 或 1 的元素作为初始阶梯。

示例 1:

输入: cost = [10, 15, 20]
输出: 15
解释: 最低花费是从cost[1]开始，然后走两步即可到阶梯顶，一共花费15。
 示例 2:

输入: cost = [1, 100, 1, 1, 1, 100, 1, 1, 100, 1]
输出: 6
解释: 最低花费方式是从cost[0]开始，逐个经过那些1，跳过cost[3]，一共花费6。
注意：

cost 的长度将会在 [2, 1000]。
每一个 cost[i] 将会是一个Integer类型，范围为 [0, 999]。

[来源](https://leetcode-cn.com/problems/min-cost-climbing-stairs/)

**思路：**
* 描述：
有一个楼梯，每次可以走1层或者2层，cost数组表示每一层所需要花费的值。可以从第一层或者第二层开始。
* 问题：
求到达顶端所花费大的最小的值。


## 代码

~~~go
package main

import (
	"fmt"
)

func main() {

	arr := []int{10, 15, 20}
	ret := minCostClimbingStairs(arr)
	fmt.Println("ret:", ret)
}

//属于动态规划问题,空间复杂度为O(n)，时间复杂度为O(n)
//[]dp 存放到达每一层需要的花费值
//cost 每一层的花费值
func minCostClimbingStairs(cost []int) int {

	n := len(cost)
	dp := make([]int, n+1) //楼顶层数

	//从0或者从1开始，所以可以得到dp[0]为0，dp[1]为0 都可以为初始值。
	//从2开始，dp[i]可以由dp[i-2]走2层或者dp[i-1]走1层，比较两者大小，选小者，其它类推
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	//最后dp[i]就是达到顶层，最优规划所花费的最小值
	fmt.Println(dp)
	return dp[n]

}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

~~~
