## 题目

给定一个链表，判断链表中是否有环。

进阶：
你能否不使用额外空间解决此题？

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
		Val: 4,
		Next: &ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  5,
					Next: nil,
				},
			},
		},
	}

	r := hasCycle(list)
	fmt.Println("ret:", r)

	//for r != nil {
	//	fmt.Println("val:", r.Val)
	//	r = r.Next
	//}

}

//快慢指针方法。如果不是环形，则快的指针先到达尾部。
func hasCycle(head *ListNode) bool {

	slow, fast := head, head.Next
	for slow != fast {

		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next

		fast = fast.Next.Next

	}

	return true
}



~~~
