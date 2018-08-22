


## 题目


给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。

你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：https://leetcode-cn.com/problems/two-sum/description/

##  代码
~~~ go

package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 3, 4, 3, 5}
	twoNums(arr, 6)

}

func twoNums(nums []int, targetValue int) []int {
	m := map[int]int{}
	idxArr := []int{}
	//转化成map
	for k, v := range nums {
		m[v] = k
	}
	fmt.Println(m)

	//从map 中查找
	for k, v := range m {
		left := targetValue - k

		if k1, ok := m[left]; ok {
			//fmt.Println("idx:", k1)

			idxArr = []int{v, k1}

			break
		}

	}
	fmt.Println("idxArr:", idxArr)

	return idxArr
}

~~~

## 扩展题目

`
在数组中找 几个数的和等于某个数。例如：找四个数的和等于12

`
**暴力解决，注意去重问题。**


## 代码
~~~go
package main

import (
	"fmt"
	"sort"
)

func main() {

	arr := []int{1, 2, 5, 4, 2, 4, 1, 7}
	r := ArrSum(arr, 12)
	fmt.Println("r:", r)
}

func IsExist(arr []int, arrCollection [][]int) bool {
	sort.Ints(arr)

	for _, v := range arrCollection {
		sort.Ints(v)
		exist := true
		for i := 0; i < len(arr); i++ {

			if arr[i] != v[i] {
				exist = false
				break
			}
		}
		if exist {
			return exist
		}

	}

	return false
}

func ArrSum(arr []int, target int) [][]int {
	r := [][]int{}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				for l := k + 1; l < len(arr); l++ {
					if arr[i]+arr[j]+arr[k]+arr[l] == target {
						currArr := []int{arr[i], arr[j], arr[k], arr[l]}
						if !IsExist(currArr, r) {
							r = append(r, []int{arr[i], arr[j], arr[k], arr[l]})
						}
					}
				}
			}
		}
	}
	return r
}


~~~
