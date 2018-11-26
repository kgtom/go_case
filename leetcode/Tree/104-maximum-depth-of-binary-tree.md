## 题目

给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

~~~
示例：
给定二叉树 [3,9,20,null,null,15,7]，

    3
   / \
  9  20
    /  \
   15   7
返回它的最大深度 3 。
~~~

[来源](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/)
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
	ret := maxDepth(node)
	fmt.Println("ret:", ret)
	fmt.Println("end")

}

//递归版：DFS 深度优先搜索
//时间复杂度 O(n)，空间复杂度：如果树是平衡的则O(logn),如果树是不平衡的，极端情况下单支树，则O(n)
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var leftValue int
	var rightValue int

	if root.Left != nil {
		leftValue = maxDepth(root.Left)
	}
	if root.Right != nil {
		rightValue = maxDepth(root.Right)
	}

	return max(leftValue, rightValue) + 1

}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}


//非递归版,迭代版 BFS:使用队列,将每一个存在左右节点的入列，然后不断迭代深入，直到最末端。
//时间复杂度 O(n)，空间复杂度 O(n)

func maxDepth2(root *TreeNode) int {

	type QueueNode struct {
		node  *TreeNode
		depth int
	}
	if root == nil {
		return 0
	}
	var maxDepth = 1
	q := []QueueNode{{node: root, depth: 1}}

	curr := q[0]
	for {

	loop:
		if curr.depth > maxDepth {
			maxDepth = curr.depth
		}
		if curr.node.Right != nil && curr.node.Left != nil {
			q = append(q, QueueNode{node: curr.node.Right, depth: curr.depth + 1})
			curr.node = curr.node.Left //改变当前链表的指向，注意：如何append 入列的是右节点，则将左节点置为当前节点。反之也可以。
		} else if curr.node.Left != nil {
			curr.node = curr.node.Left
		} else if curr.node.Right != nil {
			curr.node = curr.node.Right
		} else {
			if len(q) > 1 {
				q = q[1:]
				curr = q[0]
				goto loop //更换待深入的节点，重新开始继续深入
				fmt.Println("crr2:", curr.node.Val)
			} else {
				break
			}
		}

		curr.depth++

	}

	return maxDepth

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
