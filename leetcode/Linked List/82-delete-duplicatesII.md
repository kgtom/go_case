
## 题目

给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:
~~~
输入: 1->2->3->3->4->4->5
输出: 1->2->5

~~~
示例 2:
~~~
输入: 1->1->1->2->3
输出: 2->3
~~~

[来源](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/description/)

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
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			},
		},
	}

	cnt := ListLen(list)
	fmt.Println("len():", cnt)
	fmt.Println("--------------------")
	ret := deleteDuplicates(list)
	for ret != nil {
		fmt.Println("ret:", ret.Val)
		ret = ret.Next
	}
	fmt.Println("--------------------")
	r := deleteByVal(list, 4)

	for r != nil {
		fmt.Println("ByVal:", r.Val)

		r = r.Next
	}
	fmt.Println("--------------------")
	r2 := deleteByIdx(list, 1)

	for r2 != nil {
		fmt.Println("ByIdx:", r2.Val)
		r2 = r2.Next
	}

	fmt.Println("--------------------")
	ok := IsContain(list, 22)
	fmt.Println("isContain:", ok)

}

func deleteDuplicates(head *ListNode) *ListNode {

	curr := head
	//使用map 过滤，以node 为key,遇到重复的node，map的val +1
	var mapVal = map[int]int{}
	for curr != nil {

		mapVal[curr.Val]++

		curr = curr.Next

	}

	if head == nil {
		return head
	}
	//处理头部,即相邻不相等
	for head != nil {
		if head.Val == head.Next.Val {
			head = head.Next.Next
		} else {
			break
		}
	}

	curr = head
	repeatNode := head

	for curr != nil {
		fmt.Println("key:", curr.Val, "val:", mapVal[curr.Val], curr.Next)
		//过滤重复的node
		if mapVal[curr.Val] > 1 {
			repeatNode.Next = curr.Next
			curr = curr.Next
			continue
		}

		repeatNode = curr
		curr = curr.Next
	}
	return head
}

//根据值删除节点
func deleteByVal(head *ListNode, target int) *ListNode {

	//头部单独处理
	for head != nil {
		if head.Val == target {
			head = head.Next
		} else {
			break
		}
	}

	curr := head
	for curr.Next != nil {
		//如果相等，更改当前节点next指针，跳过next节点，直接指向next的next节点
		if curr.Next.Val == target {

			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head
}

//根据索引删除节点
func deleteByIdx(head *ListNode, idx int) *ListNode {

	if idx <= 0 {
		return head
	}
	curr := head
	i := 1
	for curr.Next != nil {

		if i == idx {

			curr.Next = curr.Next.Next

			break
		}

		curr = curr.Next
		i++
	}
	return head
}

//是否包含该节点值
func IsContain(head *ListNode, target int) bool {

	curr := head
	for curr != nil {
		if curr.Val == target {
			return true
		}
		curr = curr.Next
	}
	return false
}

//获取长度
func ListLen(head *ListNode) int {

	cnt := 0
	for head != nil {
		cnt++
		head = head.Next
	}
	return cnt
}

~~~
