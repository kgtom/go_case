## 题目

给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。



在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:
~~~
输入: 3
输出: [1,3,3,1]
~~~

进阶：

你可以优化你的算法到 O(k) 空间复杂度吗？

[来源](https://leetcode-cn.com/problems/pascals-triangle-ii/)
## 代码
~~~go
package main

import (
	"fmt"
)

func main() {

	ret := getRow(3)
	fmt.Println(ret)
}

func getRow(rowIndex int) []int {
	ret := make([]int, 1, rowIndex+1)
	ret[0] = 1 //第一个
	if rowIndex == 0 {
		return ret
	}

	for i := 0; i < rowIndex; i++ {
		ret = append(ret, 1) //每次追加1
		fmt.Println("len():", len(ret), ret)

		//当i>=2时，将重新计算索引j>=1的值
		for j := 1; j < len(ret)-1; j++ {
			ret[j] += ret[j-1]
		}
	}

	return ret
}


~~~
