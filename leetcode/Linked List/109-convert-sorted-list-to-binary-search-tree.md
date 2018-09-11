
## 题目

给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:
~~~
给定的有序链表： [-10, -3, 0, 5, 9],

一个可能的答案是：[0, -3, 9, -10, null, 5], 它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5
~~~

**思路：**因为高度不能超过1，所以保证左右两边尽量拥有相同数量的元素，那么此题就转换到二分搜索问题了，找到中间节点，左右平分。

[来源](https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/hints/)
## 代码

~~~go
package main

import "fmt"


// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	return transMidToRoot(head, nil)
}

func transMidToRoot(begin, end *ListNode) *TreeNode {
	if begin == end {
		return nil
	}

	if begin.Next == end {
		return &TreeNode{Val: begin.Val}
	}

	fast, slow := begin, begin
	for fast != end && fast.Next != end {
		fast = fast.Next.Next
		slow = slow.Next
	}

	mid := slow

	return &TreeNode{
		Val:   mid.Val,
		Left:  transMidToRoot(begin, mid),
		Right: transMidToRoot(mid.Next, end),
	}
}
func main() {

	arr := []int{-10, -3, 0, 5, 9}
	l := Slice2List(arr)
	t := sortedListToBST(l)
	fmt.Println(t)
}

func Slice2List(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	ret := &ListNode{
		Val: arr[0],
	}
	i, temp := 1, ret
	for i < len(arr) {
		temp.Next = &ListNode{
			Val: arr[i],
		}
		i++
		temp = temp.Next
	}
	return ret
}

func List2Slice(head *ListNode) []int {

	arr := []int{}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return arr
}


~~~
