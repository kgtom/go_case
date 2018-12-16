
## 题目
给定一个以字符串表示的非负整数 num，移除这个数中的 k 位数字，使得剩下的数字最小。

注意:

num 的长度小于 10002 且 ≥ k。
num 不会包含任何前导零。

示例 1 :
~~~
输入: num = "1432219", k = 3
输出: "1219"
解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。

~~~
示例 2 :

~~~
输入: num = "10200", k = 1
输出: "200"
解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。
~~~
示例 3 :
~~~
输入: num = "10", k = 2
输出: "0"
解释: 从原数字移除所有的数字，剩余为空就是0。
~~~

[来源](https://leetcode-cn.com/problems/remove-k-digits/)
## 代码
~~~go
package main

import (
	"fmt"
)

func main() {
	num := "1432219"

	ret := removeKdigits(num, 3)
	fmt.Println("ret:", ret)
}

//贪心算法：使用栈的特性，判断当前元素与栈顶元素的大小，如果小于栈顶元素，则移除栈顶元素
//time:O(n)
func removeKdigits(num string, k int) string {

	ret := []rune{}
	for _, v := range num {
		//比较当前元素v 与前一个元素大小(贪心算法的计算）
		for k > 0 && len(ret) > 0 && ret[len(ret)-1] > v {
			k--
			ret = ret[:len(ret)-1]
		}
		ret = append(ret, v)
	}
	//特殊处理，如果此刻k>0，说明ret 是递减数组，则需要去掉最后一位即可。
	for k > 0 && len(ret) > 0 {
		k--
		ret = ret[:len(ret)-1]
	}
	//特殊处理 0的情况(例如：num = "10200", k = 1)
	for len(ret) > 0 && ret[0] == '0' {
		ret = ret[1:]
	}
	// 特殊处理 空字符串
	if len(ret) == 0 {
		return "0"
	}
	return string(ret)
}

~~~
