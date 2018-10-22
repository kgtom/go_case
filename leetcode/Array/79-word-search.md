## 题目

给定一个二维网格和一个单词，找出该单词是否存在于网格中。

单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例:
~~~
board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

给定 word = "ABCCED", 返回 true.
给定 word = "SEE", 返回 true.
给定 word = "ABCB", 返回 false.
~~~
[来源](https://leetcode-cn.com/problems/word-search/description/)

## 代码
~~~go
package main

import "fmt"

func main() {

	board := [][]byte{

		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}

	ret := exist(board, "ABCCED")
	fmt.Println("ret", ret)
}
func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if DFS(board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

//深度优先遍历DFS的应用，二维数组从上下左右四个方向遍历，遇到相等返回true。
func DFS(board [][]byte, i int, j int, word string, idx int) bool {
	if idx == len(word) {
		return true
	}

	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || word[idx] != board[i][j] {
		return false
	}

	tempBoard := board[i][j]
	//遍历过的位置设置*，保证只匹配一次
	board[i][j] = '*'

	//上下左右四个方向
	isEqual := DFS(board, i+1, j, word, idx+1) || DFS(board, i, j+1, word, idx+1) || DFS(board, i-1, j, word, idx+1) || DFS(board, i, j-1, word, idx+1)

	//遍历后，再回复之前
	board[i][j] = tempBoard
	return isEqual
}



~~~
