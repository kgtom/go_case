## 题目

给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。

一般来说，删除节点可分为两个步骤：

首先找到需要删除的节点；
如果找到了，删除它。
说明： 要求算法时间复杂度为 O(h)，h 为树的高度。
~~~
示例:

root = [5,3,6,2,4,null,7]
key = 3

    5
   / \
  3   6
 / \   \
2   4   7

给定需要删除的节点值是 3，所以我们首先找到 3 这个节点，然后删除它。

一个正确的答案是 [5,4,6,2,null,null,7], 如下图所示。

    5
   / \
  4   6
 /     \
2       7

另一个正确答案是 [5,2,6,null,4,null,7]。

    5
   / \
  2   6
   \   \
    4   7
    
  ~~~

[来源](https://leetcode-cn.com/problems/delete-node-in-a-bst/)
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

	arr := []int{5, 3, 6, 2, 4, 0, 7}
	node := Int2TreeNode(arr)

	ret := deleteNode(node, 3)
	fmt.Println("ret:", ret)

}

func deleteNode(root *TreeNode, key int) *TreeNode {

	if root == nil {
		return nil
	}
	if root.Val == key {

		return mergeNode(root)
	}
	if root.Val > key {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}
	return root
}

func mergeNode(root *TreeNode) *TreeNode {

	right := root.Right
	left := root.Left

	//1.如果删除的节点没有右节点，则让它左节点代替删除的节点
	if right == nil {
		return left
	}
	node := right
	//2.如果删除的节点有右节点，且右节点下面还有左节点，则让最末层左节点替代删除的节点，即最末层node.left=left 让删除节点的左节点续上命。
	for node.Left != nil {
		node = node.Left
	}
	//3.如果删除的节点有右节点，且右节点下面没有做左节点，则让右节点替代删除节点，即右节点的左节点为删除节点的左节点
	node.Left = left
	return root.Right
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
