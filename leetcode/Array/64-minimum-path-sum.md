## 题目
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。
~~~
示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
~~~
[来源](https://leetcode-cn.com/problems/minimum-path-sum/)
## 代码
~~~go
package main

import "fmt"

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	ret := minPathSum(grid)
	fmt.Println("ret:", ret)

}

//思考：动态规划。
//求出左上角到网络中每个点的代价最小路径和，
// 若当前不是第一行，第一列，当前 grid(i,j)点，
// 那么它的值就应该是从左上角到它上面那个点 up(i-1,j)的路径和，
// 与从左上角到它左边那个点 left (i,j-1)的路径和，
// 两者中的最小值加上它自身的值。
func minPathSum(grid [][]int) int {

	if len(grid) == 0 {
		return 0
	}

	colLen := len(grid)
	rowLen := len(grid[0])

	//认识二位数组,注意边界
	//fmt.Println("grid[0][0]:", grid[0][0], grid[0][1], grid[0][2], grid[0])
	//初始化行
	//第一行的最小路径和只能从左边向右移动，所以grid[0][i] = grid[0][i] + grid[0][i-1]
	for i := 1; i < rowLen; i++ {
		grid[0][i] += grid[0][i-1]
	}
	fmt.Println("row:", grid[0]) //1 4 5
	//初始化列
	//第一列的最小路径和只能从上到下移动，所以grid[i][0] = grid[i][0] + grid[i-1][0].
	for i := 1; i < colLen; i++ {
		grid[i][0] += grid[i-1][0]
	}

	fmt.Println("col:", grid[0][0], grid[1][0], grid[2][0]) // 1 2 6

	//其它行列：比较 min()大小。即min(up,left)中小者。
	//res[i][j]=min{res[i-1][j],res[i][j-1]}+grid[i][j];
	for i := 1; i < rowLen; i++ {
		for j := 1; j < colLen; j++ {
			left := grid[i][j-1] + grid[i][j]
			up := grid[i-1][j] + grid[i][j]
			//fmt.Println("left:", left, "up:", up, grid[i][j-1], grid[i-1][j], "grid[i][j]:", grid[i][j])
			//取其小者，重塑grid[i][j]，实际：grid[i][j]+=min(up,left)
			if left > up {
				grid[i][j] = up
			} else {
				grid[i][j] = left
			}
		}
	}

	return grid[rowLen-1][colLen-1]
}

~~~
