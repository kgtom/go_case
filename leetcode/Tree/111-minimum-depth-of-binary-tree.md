## 题目
给定一个二叉树，找出其最小深度。

最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

说明: 叶子节点是指没有子节点的节点。

~~~
示例:

给定二叉树 [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回它的最小深度  2.
~~~

[来源](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/)
## 代码

~~~go
package main

import (
	"fmt"
)

// TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	arr := []int{3, 9, 20, 0, 0, 15, 7}
	node := Int2TreeNode(arr)
	ret := minDepth(node)
	fmt.Println("ret:", ret)
	fmt.Println("end")

}

//递归版：DFS 深度优先搜索
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var leftValue int
	var rightValue int

	if root.Left != nil {
		leftValue = minDepth(root.Left)
	}
	if root.Right != nil {
		rightValue = minDepth(root.Right)
	}

	return min(leftValue, rightValue) + 1

}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

//非递归版：BFS广度优先算法(一旦找到叶子节点，那么该叶子节点肯定离根节点最近，比DFS优势：不用遍历整颗树)
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
  //todo

	return 0

}
func Int2TreeNode(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: nums[0],
	}

	tempNums := make([]*TreeNode, 1)
	tempNums[0] = root

	i := 1
	for i < n {
		//每次取第一个作为根，然后再去找左右
		node := tempNums[0]
		tempNums = tempNums[1:]

		if i < n && nums[i] > 0 {
			node.Left = &TreeNode{Val: nums[i]}
			tempNums = append(tempNums, node.Left)
		}
		i++

		if i < n && nums[i] > 0 {
			node.Right = &TreeNode{Val: nums[i]}
			tempNums = append(tempNums, node.Right)
		}
		i++
	}

	return root
}

~~~
