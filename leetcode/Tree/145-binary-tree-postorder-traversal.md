## 题目
给定一个二叉树，返回它的 后序 遍历。
~~~
示例:

输入: [1,null,2,3]  
   1
    \
     2
    /
   3 

输出: [3,2,1]
~~~
[来源](https://leetcode-cn.com/problems/binary-tree-postorder-traversal/)
## 代码
~~~go
package main

import (
	"fmt"
	"sync"
)

func main() {

	//二叉树 深度遍历(前、中、后序遍历)

	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	fmt.Println("end:")
	//前序 中左右 ret: [1 2 4 5 3 6 7]
	retPreIterative := preOrderTraversal(tree)
	fmt.Println("retPreIterative:", retPreIterative)
	retPre := preOrderTraversal2(tree)
	fmt.Println("retPre-recursive:", retPre)

	fmt.Println()
	//中序 迭代版 左中右 ret: [4 2 5 1 6 3 7]
	retInIterative := inOrderTraversal(tree)
	fmt.Println("retInIterative", retInIterative)
	//中序 递归版
	retIn := inOrderTraversal2(tree)
	fmt.Println("retIn-recursive:", retIn)

	fmt.Println()

	//后序 左右中 ret: [1 2 4 5 3 6 7]
	retPostIterative := postOrderTraverse(tree)
	fmt.Println("retPostIterative:", retPostIterative)
	//后序 迭代版
	retPost := postOrderTraverse2(tree)
	fmt.Println("retPost-recursive:", retPost)
	//层序=前序 ret: [1 2 4 5 3 6 7]

	fmt.Println("end:")

}

// TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack struct {
	lock sync.Mutex
	node []*TreeNode
}

func NewStack() *Stack {

	return &Stack{lock: sync.Mutex{}, node: []*TreeNode{}}
}

func (s *Stack) Push(node *TreeNode) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.node = append(s.node, node)
}

func (s *Stack) Pop() *TreeNode {
	s.lock.Lock()

	defer s.lock.Unlock()
	n := len(s.node)
	if n == 0 {
		return nil
	}
	ret := s.node[n-1]
	s.node = s.node[:(n - 1)]

	return ret
}

//前序 迭代版：(中左右)时间复杂度O(n) 空间复杂度(n)
func preOrderTraversal(tree *TreeNode) []int {

	s := NewStack()

	//将tree,push 到栈里面
	s.Push(tree)
	ret := []int{}

	//使用stack 栈，注意其特性，先进后出。

	for len(s.node) != 0 {

		//获取当前节点,首次实际就是根节点值
		curr := s.Pop()
		ret = append(ret, curr.Val)

		//如果当前节点存在右节点，则入栈
		//如果当前节点存在左节点，则入栈
		//先 push right节点，后 push left,是因为栈的特性：先进后出，又因为前序遍历，需要先让左节点出来，后右节点出来。
		if curr.Right != nil {
			s.Push(curr.Right)
		}
		if curr.Left != nil {
			s.Push(curr.Left)
		}

	}
	//fmt.Println("len:", len(s.node)) //0
	return ret
}

//前序 递归版
func preOrderTraversal2(root *TreeNode) []int {
	var ret []int
	if root == nil {
		return ret
	}
	//fmt.Println(root.Val)
	ret = append(ret, root.Val)
	ret = append(ret, preOrderTraversal2(root.Left)...)
	ret = append(ret, preOrderTraversal2(root.Right)...)

	return ret
}

//  中序 迭代版 ：(左中右)时间复杂度O(n) 空间复杂度(n)
func inOrderTraversal(root *TreeNode) []int {
	var ret []int

	if root == nil {
		return ret
	}

	stack := NewStack()
	curr := root

	for {
		//现将root入栈，迭代找到左节点的起点，即左节点为nil的节点。
		for curr != nil {
			stack.Push(curr)
			curr = curr.Left
		}

		//因为上一步左节点最后一个入栈，则第一个出栈，符合中序遍历，先左节点开始，然后中节点，最后是右节点
		node := stack.Pop()
		if node == nil {
			break
		}

		ret = append(ret, node.Val)

		// 再将右节点入栈，入栈是 右节点部分也要按照 左-中-右
		if node.Right != nil {
			curr = node.Right
		}
	}

	return ret
}

//中序 递归版
func inOrderTraversal2(root *TreeNode) []int {
	var ret []int

	if root == nil {
		return ret
	}

	ret = append(ret, inOrderTraversal2(root.Left)...)
	ret = append(ret, root.Val)
	ret = append(ret, inOrderTraversal2(root.Right)...)

	return ret
}

//后序 迭代版 左右中
func postOrderTraverse(root *TreeNode) []int {
	var ret []int

	if root == nil {
		return ret
	}
	stemp := NewStack()
	s := NewStack()

	stemp.Push(root)

	for {

		node := stemp.Pop()

		if node == nil {
			break
		}
		s.Push(node)

		if node.Left != nil {

			stemp.Push(node.Left)

		}
		if node.Right != nil {
			stemp.Push(node.Right)
		}

	}
	for {
		node := s.Pop()
		if node == nil {
			break
		}
		ret = append(ret, node.Val)
	}

	return ret
}

//后序 递归版
func postOrderTraverse2(root *TreeNode) []int {
	var ret []int
	if root == nil {
		return ret
	}

	ret = append(ret, postOrderTraverse2(root.Left)...)

	ret = append(ret, postOrderTraverse2(root.Right)...)
	ret = append(ret, root.Val)
	//fmt.Println(root.Val)

	return ret
}

~~~
