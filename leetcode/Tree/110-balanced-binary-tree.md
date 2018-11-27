## 题目

给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
~~~
示例 1:

输入:
    2
   / \
  1   3
输出: true
示例 2:

输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。
     
~~~

[来源](https://leetcode-cn.com/problems/validate-binary-search-tree/)

## 代码

~~~go
package main

import (
	"fmt"
	"math"
)

// TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	arr := []int{2, 1, 3}
	node := Int2TreeNode(arr)

	ret := isValidBST2(node)
	fmt.Println("ret:", ret)

}

// 递归版--网上大神的做法
//对于某个节点n来说：
//n的左子树的范围一定是(min,n.val)
//n的右子树的范围一定是(n.val,max)
//左右节点必满足： n.val>min &&n.val<max,否则return false
func isValidBST(n *TreeNode) bool {
	if n == nil || (n.Left == nil && n.Right == nil) {
		return true
	}
	return isValidTree(n, math.MaxInt64, math.MinInt64)
}

func isValidTree(n *TreeNode, max int, min int) bool {
	if n == nil {
		return true
	}
	if n.Val > min && n.Val < max {
		return isValidTree(n.Left, n.Val, min) && isValidTree(n.Right, max, n.Val)
	}
	return false
}

// 递归版--中序遍历
//根据二叉查找树的定义，将其中序遍历可得到递增序列，判断序列大小就可以验证是否是二叉查找树
func isValidBST2(n *TreeNode) bool {

	arr := inOrderTraversal(n)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

//中序遍历递归版
func inOrderTraversal(root *TreeNode) []int {
	var ret []int

	if root == nil {
		return ret
	}

	ret = append(ret, inOrderTraversal(root.Left)...)
	ret = append(ret, root.Val)
	ret = append(ret, inOrderTraversal(root.Right)...)

	return ret
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

 
ps: [大神算法地址](https://leetcode.com/problems/validate-binary-search-tree/discuss/179030/Beats-100-Golang)
