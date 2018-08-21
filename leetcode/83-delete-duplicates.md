## 题目
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:
~~~
输入: 1->1->2
输出: 1->2
~~~

示例 2:
~~~
输入: 1->1->2->3->3
输出: 1->2->3
~~~

[来源](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/)

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
			Val: 1,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}
	r := deleteDuplicates(list)
	//遍历
	for r != nil {
		fmt.Println("r:", r.Val)
		fmt.Println("--------------------")
		r = r.Next
	}

}

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head
	for curr != nil && curr.Next != nil {

		//比较相邻节点，若相等，则跳过下一个节点，将next节点赋值为next的下一个节点
		if curr.Next.Val == curr.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}


~~~
