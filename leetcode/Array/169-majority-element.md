## 题目

给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在众数。

示例 1:
~~~
输入: [3,2,3]
输出: 3

~~~
示例 2:
~~~
输入: [2,2,1,1,1,2,2]
输出: 2
~~~
[来源](https://leetcode-cn.com/problems/majority-element/)
## 代码

~~~go
package main

import (
	"fmt"
)

func main() {

	ret := majorityElement2([]int{3, 5, 3})
	fmt.Println("ret:", ret)

}
func majorityElement(nums []int) int {
	//巧用map,使用数组元素做map的key,然后⌊ n/2 ⌋与比较
	m := make(map[int]int)
	fmt.Println(len(nums), len(nums)/2)
	for _, v := range nums {
		m[v]++
		//出现次数大于 ⌊ n/2 ⌋ 的元素
		if m[v] > len(nums)/2 {
			return v
		}
	}
	fmt.Println(m)
	return 0
}

func majorityElement2(nums []int) int {

	//从默认第一个开始，遇到相同的count++，遇到不同的count--，如果count==0，则另一换一个defaultVal
	defaultVal, count := len(nums), 1
	for _, v := range nums {
		//if defaultVal == v {
		//	count++
		//} else if count > 0 {
		//	count--
		//} else {
		//	defaultVal = v
		//}
		//ps 也可以使用switch-case替换if-else if -else

		switch {
		case defaultVal == v:
			count++
		case count > 0:
			count--
		default:
			defaultVal = v
		}
	}
	return defaultVal
}

~~~
