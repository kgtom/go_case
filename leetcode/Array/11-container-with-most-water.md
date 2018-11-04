
## 题目

给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器，且 n 的值至少为 2。



图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

 
~~~
示例:

输入: [1,8,6,2,5,4,8,3,7]
输出: 49
~~~

[来源](https://leetcode-cn.com/problems/container-with-most-water/)
## 代码
~~~go
package main

import (
	"fmt"
	"math"
)

func main() {

	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	ret := maxArea2(arr)
	fmt.Println("ret:", ret)
}

//1.双指针方法 ，一次遍历，时间复杂度O(n) 空间复杂度O(1)使用恒定空间
func maxArea(height []int) int {

	l, r := 0, len(height)-1
	ret, tempArea := 0, 0
	for l < r {

		// height[l] 和 height[r] 代表矩形宽
		//r-l 代表矩形长
		//最大值 取决于较短的宽，实际木桶效应，盛水多少取决于短的
		if height[l]*(r-l) > height[r]*(r-l) {
			tempArea = height[r] * (r - l)
		} else {
			tempArea = height[l] * (r - l)
		}
		fmt.Println("l:", l, "r:", r, "height:", height[l], "width:", r-l, "area:", tempArea)
		if tempArea > ret {
			ret = tempArea
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}

	}
	return ret
}

//2.暴力方法 遍历 时间复杂度O(n2) 空间复杂度O(1)恒定

func maxArea2(height []int) int {

	ret, temp := 0, 0

	for i := 0; i < len(height)-1; i++ {
		for j := len(height) - 1; j >= 0; j-- {
			if height[i] > height[j] {
				temp = height[j] * (j - i)
			} else {
				temp = height[i] * (j - i)
			}
			if temp > ret {
				ret = temp
			}
		}
	}
	return ret
}

//3.暴力方法 使用 math.Max
func maxArea3(height []int) int {

	var ret float64

	for i := 0; i < len(height)-1; i++ {
		for j := len(height) - 1; j >= 0; j-- {
			if height[i] > height[j] {

				ret = math.Max(ret, float64(height[j]*(j-i)))
			} else {

				ret = math.Max(ret, float64(height[i]*(j-i)))
			}

		}
	}
	return int(ret)
}


~~~
