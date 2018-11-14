## 题目

给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地) 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。

找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)
~~~
示例 1:

[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
对于上面这个给定矩阵应返回 6。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。

示例 2:

[[0,0,0,0,0,0,0,0]]
对于上面这个给定的矩阵, 返回 0。
~~~

注意: 给定的矩阵grid 的长度和宽度都不超过 50。

[来源](https://leetcode-cn.com/problems/max-area-of-island/)
## 代码
~~~go
package main

import "fmt"

func main() {

	grid := [][]int{[]int{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		[]int{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		[]int{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}

	ret := maxAreaOfIsland(grid)
	fmt.Println("ret", ret)
}

//深度优先算法。
// 注意两点： 1.找到着陆点即非0，然后上下左右开始遍历，遇到0返回，得到area，多个area,返回最大的即可
// 2.对已经遍历的坐标置零，避免重复统计
func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			//从grid[row][colum]==1 开始
			if grid[i][j] == 1 {
				//fmt.Println("start row:", i, "colum:", j)
				area := getArea2(grid, i, j)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	return maxArea
}

//第一种写法，分开遍历
func getArea(grid [][]int, row, colum int) int {
	if grid[row][colum] == 0 {
		return 0
	}

	//递归时注意条件：即 从grid[row][colum]==1 开始
	area := 1
	// 当前元素设置为0，避免再次搜到
	grid[row][colum] = 0

	if row > 0 {

		area += getArea(grid, row-1, colum)

	}
	if row < len(grid)-1 {

		area += getArea(grid, row+1, colum)

	}
	if colum > 0 {

		area += getArea(grid, row, colum-1)

	}
	if colum < len(grid[0])-1 {

		area += getArea(grid, row, colum+1)

	}

	return area
}

//第二种写法，一起遍历
func getArea2(grid [][]int, row, colum int) int {

	if grid[row][colum] == 0 {
		return 0
	}

	//递归时注意条件：即 从grid[row][colum]==1 开始
	area := 1
	// 当前元素设置为0，避免再次搜到
	grid[row][colum] = 0

	if row > 0 && row < len(grid)-1 && colum > 0 && colum < len(grid[0])-1 {

		area += getArea(grid, row-1, colum) + getArea(grid, row+1, colum) + getArea(grid, row, colum-1) + getArea(grid, row, colum+1)

	}

	return area
}

~~~
