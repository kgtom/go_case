## 题目

给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

示例:
~~~
给定 1->2->3->4, 你应该返回 2->1->4->3.
~~~
说明:

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

**思路：** 递归，每次递归返回值组合

## 代码

~~~go
package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	newHead := head.Next
	fmt.Println("newHead:", newHead)
	head.Next = swapPairs(newHead.Next)
	fmt.Println("head.next:", head.Next)
	fmt.Println("head:", head)
	newHead.Next = head
	fmt.Println("newHead:", newHead)
	return newHead
}

func main() {

	arr := []int{1, 2, 3, 4}
	l := Slice2List(arr)
	r := swapPairs(l)
	for r != nil {
		fmt.Println("val:", r.Val)
		r = r.Next
	}

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
