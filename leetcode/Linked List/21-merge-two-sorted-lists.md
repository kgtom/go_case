
## 题目
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

示例：
~~~
输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
~~~
[来源](https://leetcode-cn.com/problems/merge-two-sorted-lists/description/)
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
			Val: 3,
			Next: &ListNode{
				Val:  5,
				Next: nil,
			},
		},
	}

	list2 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}

	r := mergeTwoLists(list, list2)

	for r != nil {
		fmt.Println("val:", r.Val)
		r = r.Next
	}

}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	curr := &ListNode{}
	ret := curr
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			curr.Next = l2
			l2 = l2.Next
		} else {
			curr.Next = l1
			l1 = l1.Next
		}
		curr = curr.Next

	}
	if l1 == nil {
		curr.Next = l2
	}
	if l2 == nil {
		curr.Next = l1
	}

	return ret.Next
}

~~~
