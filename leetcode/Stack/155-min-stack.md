## 题目

设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。

push(x) -- 将元素 x 推入栈中。
pop() -- 删除栈顶的元素。
top() -- 获取栈顶元素。
getMin() -- 检索栈中的最小元素。

~~~
示例:

MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.getMin();   --> 返回 -2.

~~~

[来源](https://leetcode-cn.com/problems/min-stack/)

## 代码

~~~go
package main

import "fmt"

func main() {

	minStack := Constructor()

	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println("min:", minStack.GetMin()) // 返回 -3
	minStack.Pop()
	fmt.Println("top:", minStack.Top())    //返回0
	fmt.Println("min:", minStack.GetMin()) // 返回 -2

	
}

type MinStack struct {
	nums []int
}

/** initialize your data structure here. */
func Constructor() MinStack {

	return MinStack{nums: []int{}}
}

func (this *MinStack) Push(x int) {
	this.nums = append(this.nums, x)
}

func (this *MinStack) Pop() {

	//ret := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]

}

//栈顶元素，最后push到数组的
func (this *MinStack) Top() int {
	return this.nums[len(this.nums)-1]
}

func (this *MinStack) GetMin() int {

	ret := 0
	for i := 0; i < len(this.nums); i++ {
		if ret > this.nums[i] {
			ret = this.nums[i]
		}
	}
	return ret
}



~~~
