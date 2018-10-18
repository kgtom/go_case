
## 题目 
对链表进行插入排序。
![Insertion-sort-example-300px](/content/images/2018/10/Insertion-sort-example-300px.gif)

插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。

 

插入排序算法：

插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
重复直到所有输入数据插入完为止。
 

示例 1：
~~~
输入: 4->2->1->3
输出: 1->2->3->4
~~~

示例 2：
~~~
输入: -1->5->3->4->0
输出: -1->0->3->4->5
~~~

[来源](https://leetcode-cn.com/problems/insertion-sort-list/description/)

## 代码

~~~
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	list := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val:  3,
					Next: nil,
				},
			},
		},
	}
	ret := insertionSortList(list)
	for ret != nil {
		fmt.Println("val:", ret.Val)
		ret = ret.Next
	}

}

//时间复杂度O(n^2)，效率不高的算法，但空间复杂度O(1),以高时间复杂度换取低空间复杂度
func insertionSortList(head *ListNode) *ListNode {
	ret := &ListNode{} //新链表
	curr := head

	for curr != nil {

		pre := ret
		//比较新链表大小排序
		for pre.Next != nil && pre.Next.Val < curr.Val {

			pre = pre.Next

		}
		//迭代当前链表
		next := curr.Next
		curr.Next = pre.Next
		pre.Next = curr
		curr = next
	}
	return ret.Next
}


~~~
