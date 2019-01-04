## 题目

给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。

说明：
你可以假设 k 总是有效的，1 ≤ k ≤ 二叉搜索树元素个数。

示例 1:
~~~
输入: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2
输出: 1
~~~
示例 2:
~~~
输入: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1
输出: 3
~~~
进阶：
如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化 kthSmallest 函数？




## 代码
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 1,
			Right: &TreeNode{
				Val: 2,
			},
		},
		Right: &TreeNode{
			Val: 4,
		},
	}
	ret := kthSmallest(tree, 1)
	fmt.Println("ret:", ret)

}

//首先理解二叉搜索树特性，左边小于右边值，然后使用二分查找，定位最小值在左还是右进行查找,定位关键在于k 与左边节点个数比较
func kthSmallest(root *TreeNode, k int) int {

	//获取左边最小节点个数
	minCount := getTreeCount(root.Left)
	//说明是根节点
	if minCount+1 == k {
		return root.Val
	} else if minCount < k {
		//右边查找
		return kthSmallest(root.Right, k-minCount-1)
	} else {
		//左边查找
		return kthSmallest(root.Left, k)
	}
	return k
}

//递归获得节点个数
func getTreeCount(root *TreeNode) int {

	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	} else {
		left := getTreeCount(root.Left)
		right := getTreeCount(root.Right)
		return 1 + left + right
	}
}


~~~
