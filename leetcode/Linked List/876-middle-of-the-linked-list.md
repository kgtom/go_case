## 题目
给定一个带有头结点 head 的非空单链表，返回链表的中间结点。

如果有两个中间结点，则返回第二个中间结点。

 

示例 1：
~~~
输入：[1,2,3,4,5]
输出：此列表中的结点 3 (序列化形式：[3,4,5])
返回的结点值为 3 。 (测评系统对该结点序列化表述是 [3,4,5])。
注意，我们返回了一个 ListNode 类型的对象 ans，这样：
ans.val = 3, ans.next.val = 4, ans.next.next.val = 5, 以及 ans.next.next.next = NULL.
~~~
示例 2：
~~~
输入：[1,2,3,4,5,6]
输出：此列表中的结点 4 (序列化形式：[4,5,6])
由于该列表有两个中间结点，值分别为 3 和 4，我们返回第二个结点。
~~~
[来源](https://leetcode-cn.com/problems/middle-of-the-linked-list/submissions/1)

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

	r := middleNode(list)
	fmt.Println("ret:", r)

	fmt.Println("-------------------")
	for list != nil {
		fmt.Println("val:", list.Val)
		list = list.Next
	}
	fmt.Println("end")
}

//输出数组方法： 将节点值放在数组中，通过len/2索引找到中间节点
func middleNode(head *ListNode) int {

	var arr []int

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	if len(arr) == 0 {
		return 0
	}
	return arr[len(arr)/2]

}

// 遍历节点找到中间节点
func middleNode2(head *ListNode) int {

	middleIdx := listLen(head) / 2
	i := 0
	for head != nil {

		if i == middleIdx {
			return head.Val
		}
		i++
		head = head.Next

	}
	return i
}

func listLen(head *ListNode) int {

	cnt := 0
	for head != nil {
		cnt++
		head = head.Next
	}

	return cnt
}

~~~
