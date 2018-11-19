
## 题目
给定一个二叉树，返回它的 前序 遍历。
~~~
 示例:

输入: [1,null,2,3]  
   1
    \
     2
    /
   3 

输出: [1,2,3]
~~~
进阶: 递归算法很简单，你可以通过迭代算法完成吗？



## 代码

### 迭代实现思路：

根据stack的特性：后入先出，所以我们在入栈Push的时候，先将右子树入栈，再将左子树入栈，这样在出栈Pop的时候就可以保证是先左后右，符合前序遍历(中左右)。

~~~go
package main

import (
	"fmt"
	"sync"
)

func main() {

	//将二叉树按照前序遍历输出到[]int中

	tree := &TreeNode{
		Val: 1,
		//Left: &TreeNode{
		//	Val: 4,
		//},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	ret := preorderTraversal(tree)
	fmt.Println("ret:", ret)

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

//二叉树前序遍历：时间复杂度O(n) 空间复杂度(n)
func preorderTraversal(tree *TreeNode) []int {

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
		//先给right节点，后给left,是因为栈的特性：先进后出，又因为前序遍历，需要先让左节点出来，后右节点出来。
		if curr.Right != nil {
			s.Push(curr.Right)
		}
		if curr.Left != nil {
			s.Push(curr.Left)
		}

	}
	fmt.Println("len:", len(s.node)) //0
	return ret
}

~~~
