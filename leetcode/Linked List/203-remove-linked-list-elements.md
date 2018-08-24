## 题目
删除链表中等于给定值 val 的所有节点。

示例:
~~~
输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
~~~
[来源](https://leetcode-cn.com/problems/remove-linked-list-elements/description/)
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
				Val: 6,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  6,
						Next: nil,
					},
				},
			},
		},
	}

	r := removeElements(list, 6)

	for r != nil {
		fmt.Println("val:", r.Val)
		r = r.Next
	}

}

func removeElements(head *ListNode, targetVal int) *ListNode {

	curr := head
	if curr == nil {
		return curr
	}
	//单独处理头部
	if curr.Val == targetVal {
		head = head.Next
		return head
	}

	//除头部之外
	for curr.Next != nil {
		if curr.Next.Val == targetVal {
			curr.Next = curr.Next.Next
			//return head 删除一个，去掉注释则删除所有的targetVal
		} else {

			curr = curr.Next
		}

	}
	return head
}

~~~
