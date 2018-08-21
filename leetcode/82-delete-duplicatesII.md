
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
	//使用map 过滤，以node 为key,遇到重复的node，map的val +1
	var mapVal = map[int]int{}
	for curr != nil {

		mapVal[curr.Val]++

		curr = curr.Next

	}

	//查看map 数据
	for k, v := range mapVal {
		fmt.Println(k, v)
		//3 2
		//1 2
		//2 1
	}

	retNode := head
	for curr != nil {

		//过滤重复的node
		if mapVal[curr.Val] > 1 {
			retNode.Next = curr.Next
			curr = curr.Next
			continue
		}

		retNode = curr
		curr = curr.Next
	}
	return retNode
}

~~~
