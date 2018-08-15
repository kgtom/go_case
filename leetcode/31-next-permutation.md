## 题目
`
实现下一个排列，它将数字重新排列成字典下一个更大的数字排列。

如果这种安排不可能，则必须将其重新排列为尽可能低的顺序（即按升序排序）。

更换必须就地，并且只使用恒定的额外内存。

这里有些例子。输入位于左侧列中，其相应的输出位于右侧列中。

1,2,3→1,3,2
3,2,1→1,2,3
1,1,5→1,5,1

`

来源：https://leetcode.com/problems/next-permutation/description/

## 代码
~~~go
package main

import (
	"fmt"
	"sort"
)

func main() {

	//arr := []int{1, 3, 4, 5, 9} //[1 3 4 9 5]
	//arr := []int{9, 5, 4, 3, 1} //[1 3 4 9 5]
	//arr := []int{9, 5, 1, 3, 4} //[9 5 1 4 3]
	arr := []int{1, 3, 9, 5, 4} //1 4 3 5 9
	nextPermutation(arr)
	fmt.Println("r:", arr)
}
func nextPermutation(arr []int) {
	isSeq := true
	for i := len(arr) - 1; i > 0; i-- {

		if arr[i] > arr[i-1] {
			isSeq = false
			swapIdx := i
			for j := i + 1; j < len(arr); j++ {
				//拿着当前者，左邻居与其右邻居比较，找到右邻居中最小的交换位置
				if arr[j] > arr[i-1] && arr[j] < arr[i] {
					swapIdx = j
				}
			}
			arr[i-1], arr[swapIdx] = arr[swapIdx], arr[i-1]
			fmt.Println(arr)
			sort.Ints(arr[i:]) //包含自己升序
			fmt.Println("re:", arr)
			break
		}
	}

	if isSeq {
		k := len(arr) / 2
		//fmt.Println("k:", k)
		for i := 0; i < k; i++ {
			fmt.Println(arr[i], arr[len(arr)-1-i])
			arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
		}
	}
}
~~~
