## 题目

给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

示例：

输入：(2 -> 4 -> 3) + (4 -> 6 -> 4)
输出：6 -> 0 -> 8
原因：242 + 465 = 608

来源：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/
 


## 代码

~~~go
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	l2 := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	r := addTwoNums(l1, l2)
	fmt.Println(r.Val, r.Next.Val, r.Next.Next.Val)
}

func addTwoNums(l1 *ListNode, l2 *ListNode) *ListNode {
	var res = &ListNode{}
	curr := res
	for {
		if l1 != nil {
			curr.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			curr.Val += l2.Val
			l2 = l2.Next
		}
		//迭代
		curr.Next = &ListNode{}
		curr = curr.Next

		if l1 == nil && l2 == nil {
			break
		}
	}

	// 进位
	fmt.Println("curr 2:", curr.Val, curr, "res:", res.Val, res)
	curr2 := res
	fmt.Println("curr 3:", curr2)
	for curr2 != nil {
		if curr2.Val >= 10 {
			if curr2.Next == nil {
				curr2.Next = &ListNode{}
			}

			curr2.Val = curr2.Val - 10
			curr2.Next.Val++
		}

		//迭代
		curr2 = curr2.Next
	}

	return res
}

~~~
