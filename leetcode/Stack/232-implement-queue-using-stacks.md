## 题目
使用栈实现队列的下列操作：

push(x) -- 将一个元素放入队列的尾部。
pop() -- 从队列首部移除元素。
peek() -- 返回队列首部的元素。
empty() -- 返回队列是否为空。

~~~
示例:

MyQueue queue = new MyQueue();

queue.push(1);
queue.push(2);  
queue.peek();  // 返回 1
queue.pop();   // 返回 1
queue.empty(); // 返回 false

~~~
说明:

你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size, 和 is empty 操作是合法的。
你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。

[来源](https://leetcode-cn.com/problems/implement-queue-using-stacks/)
## 代码
~~~go
package main

import (
	"fmt"
)

func main() {

	queue := Constructor()

	queue.Push(1)
	queue.Push(2)
	fmt.Println("peek:", queue.Peek()) // 返回 1
	fmt.Println("pop:", queue.Pop())   // 返回 1

	fmt.Println("empty:", queue.Empty()) // 返回 false
}

// 使用两个栈，a 入栈，b是a的出栈。 记住 栈：先入后出，队列 先入先出。
type MyQueue struct {
	a, b *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {

	return MyQueue{a: NewStack(), b: NewStack()}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {

	this.a.Push(x)

}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {

	if len(this.b.num) == 0 {

		for len(this.a.num) > 0 {
			this.b.Push(this.a.Pop())
		}
	}

	return this.b.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.b.num) == 0 {

		for len(this.a.num) > 0 {
			this.b.Push(this.a.Pop())
		}
	}
	//获取栈 b的第一个元素，即栈顶的元素，也就是最后push进去的元素，在数组中的索引是最后一位。
	return this.b.num[len(this.b.num)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {

	return len(this.b.num) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

type Stack struct {
	num []int
}

func NewStack() *Stack {

	return &Stack{num: []int{}}
}
func (s *Stack) Push(i int) {

	s.num = append(s.num, i)

}
func (s *Stack) Pop() int {

	ret := s.num[len(s.num)-1]
	s.num = s.num[:(len(s.num) - 1)]

	return ret
}

~~~
