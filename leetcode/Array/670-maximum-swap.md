## 题目
给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。

~~~
示例 1 :

输入: 2736
输出: 7236
解释: 交换数字2和数字7。
示例 2 :

输入: 9973
输出: 9973
解释: 不需要交换。

~~~
注意:
给定数字的范围是 [0, 108]

[来源](https://leetcode-cn.com/problems/maximum-swap/description/)

## 代码

~~~go

package main

import (
	"fmt"
	"strconv"
)

func main() {

	num := 97387

	ret := maximumSwap(num)
	fmt.Println("nums2:", ret)

}
func maximumSwap(num int) int {

	bs := []byte(strconv.Itoa(num))

	// buckets桶 记录了每个数字最后出现的位置
	buckets := make(map[byte]int, len(bs))
	for i := range bs {
		buckets[bs[i]] = i
	}

	// 遍历bs,与buckets 桶中比较大小，桶里数字大的，两者交换
	for i := 0; i < len(bs); i++ {
		//最大数字9开始遍历,bs与桶里数比较
		for bj := byte('9'); bs[i] < bj; bj-- {
			k := buckets[bj]
			//k位置大于i,因为前面位数已有数字，不需要小于i位置的数字
			if k > i {
				bs[i], bs[k] = bs[k], bs[i]
				return convert2Int(bs)
			}
		}
	}

	return convert2Int(bs)
}

func convert2Int(bs []byte) int {
	n, _ := strconv.Atoi(string(bs))
	return n
}

~~~
