
## 题目


给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。

示例 1:

~~~
给定数组 nums = [1,1,2], 

函数应该返回新的长度 2, 并且原数组 nums 的前两个元素被修改为 1, 2。 

你不需要考虑数组中超出新长度后面的元素。

~~~

示例 2:

~~~
给定 nums = [0,0,1,1,1,2,2,3,3,4],

函数应该返回新的长度 5, 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4。

你不需要考虑数组中超出新长度后面的元素。

~~~

说明:

为什么返回数值是整数，但输出的答案是数组呢?

请注意，输入数组是以“引用”方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。

你可以想象内部操作如下:

~~~
// nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
int len = removeDuplicates(nums);

// 在函数里修改输入数组对于调用者是可见的。
// 根据你的函数返回的长度, 它会打印出数组中该长度范围内的所有元素。
for (int i = 0; i < len; i++) {
    print(nums[i]);
}
~~~

[来源：](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/description/)

## 代码

~~~go
package main

import "fmt"

func main() {
	//arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4} //[0,0,1,1,1,2,2,3,3,4]
	arr := []int{1, 2, 2, 3, 3} //1, 2, 2, 3, 3
	r := removeDuplicates2(arr)
	fmt.Println("r:", r)
}

//记录不同元素的个数
func removeDuplicates(arr []int) int {

	cnt := 0

	for i := 1; i < len(arr); i++ {
		if arr[cnt] != arr[i] {
			cnt++
			arr[cnt] = arr[i]

		}
		//else {
		//	fmt.Println(arr[cnt], arr[i])
		//	continue
		//}
	}
	fmt.Println("arr:", arr[:cnt+1])
	return cnt + 1
}

//相邻元素相同的继续迭代，不同的元素记录次数并赋值前移
func removeDuplicates2(arr []int) int {

	cnt := 0
	i := 1

	for i < len(arr) {
		if arr[i] == arr[cnt] {
			i++

		} else {
			cnt++
			arr[cnt] = arr[i]
			//fmt.Println(arr[cnt], arr[i])
			i++
		}

	}
	fmt.Println("arr:", arr[:cnt+1], "src arr:", arr)
	return cnt + 1
}

~~~
