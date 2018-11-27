## 题目


给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

~~~
示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4
返回 false 。

~~~

[来源](https://leetcode-cn.com/problems/balanced-binary-tree/)


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

	arr := []int{3, 9, 20, 0, 0, 15, 76}
	node := Int2TreeNode(arr)

	ret := isBalanced(node)
	fmt.Println("ret:", ret)

}

// 递归版：时间复杂度O(n)
func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}

	leftDepth, rightDepth := getDepth(root.Left), getDepth(root.Right)
	//高度差的大于1，则不平衡
	if leftDepth > rightDepth+1 || rightDepth > leftDepth+1 {
		return false
	}
	return true
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

func getDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var leftValue int
	var rightValue int
	//fmt.Println("root:", root.Val)
	if root.Left != nil {

		leftValue = getDepth(root.Left)
	}
	if root.Right != nil {
		rightValue = getDepth(root.Right)
	}

	//返回当前分支树的最大深度，左右取大者。
	return max(leftValue, rightValue) + 1

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}


~~~
