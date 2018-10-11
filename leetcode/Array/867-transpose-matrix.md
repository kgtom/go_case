## 题目
矩阵的转置是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。

 

示例 1：
~~~
输入：[[1,2,3],[4,5,6],[7,8,9]]
输出：[[1,4,7],[2,5,8],[3,6,9]]
~~~

示例 2：
~~~
输入：[[1,2,3],[4,5,6]]
输出：[[1,4],[2,5],[3,6]]
 ~~~

提示：

1 <= A.length <= 1000
1 <= A[0].length <= 1000

[来源](https://leetcode-cn.com/problems/transpose-matrix/)
## 代码

~~~go
package main

import (
	"fmt"
)

func main() {

	ret := transpose([][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}})
	fmt.Println("ret:", ret)

}
func transpose(nums [][]int) [][]int {

	l, col := len(nums), len(nums[0]) //行，列
	ret := make([][]int, l)
	for i := range ret {
		ret[i] = make([]int, l)

	}
	for i := 0; i < l; i++ {
		for j := 0; j < col; j++ {
			ret[j][i] = nums[i][j] //观察演示数据得知，需要交换矩阵的行索引与列索引即可。
		}
	}

	return ret
}


~~~
