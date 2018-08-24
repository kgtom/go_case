## 题目

反转一个单链表。

示例:
~~~
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
~~~

进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

[来源](https://leetcode-cn.com/problems/reverse-linked-list/description/)
## 代码
~~~go
package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	list := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 33,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	//三种都可以
	r := reverseList(list)
	//r := reverseList2(list)
	//r := reverseList3(list, nil)
	for r != nil {
		fmt.Println("val:", r.Val)
		r = r.Next
	}

}

//迭代方法：得到下一个节点，更改当前节点使指针下移。
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev

}

//递归法: 1 2 3 ，把 2放在1的前面,1的后面是nil

func reverseList2(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}
	curr := head
	return reverseList3(curr, nil)
}
func reverseList3(curr *ListNode, prev *ListNode) *ListNode {

	if curr == nil {
		return prev
	}
	next := curr.Next
	curr.Next = prev
	prev = curr //这一句不能与 curr = next 换执行顺序。
	curr = next

	return reverseList3(curr, prev)

}

~~~
