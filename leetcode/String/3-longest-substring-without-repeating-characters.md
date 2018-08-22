## 题目

~~~
给定一个字符串，找出不含有重复字符的最长子串的长度。

示例：

给定 "abcabcbb" ，没有重复字符的最长子串是 "abc" ，那么长度就是3。

给定 "bbbbb" ，最长的子串就是 "b" ，长度是1。

给定 "pwwkew" ，最长子串是 "wke" ，长度是3。请注意答案必须是一个子串，"pwke" 是 子序列  而不是子串。
~~~

来源:https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/description/

## 代码

~~~go
package main

import "fmt"

func main() {
	s := "abcabcd"
	r := lengthfLongestSubstring(s)
	fmt.Println("r:", r)
}

//滑动窗口原理
func lengthOfLongestSubstring(s string) int {
	var arr [128]int
	var i, maxLen int
	for k, v := range s {
		fmt.Println(arr[v], k, v)
		//fmt.Println("arr[v]:", arr[v], "i:", i)
		i = max(arr[v], i)
		//fmt.Println("maxLen:", maxLen, "max-k:", k-i+1)
		//fmt.Println("-----------------")
		maxLen = max(maxLen, k-i+1)
		arr[v] = k + 1
	}
	fmt.Println(arr)
	return maxLen
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

//使用map,将重复的删除
func lengthLongestSubstring(s string) int {
	var checkMap = map[uint8]bool{}
	var maxLen = 0
	var currLen = 0

	for i := 0; i < len(s); i++ {
		if checkMap[s[i]] == true {

			delIdx := i - currLen

			delete(checkMap, s[delIdx])
			currLen--

		}
		checkMap[s[i]] = true
		currLen++
		if currLen > maxLen {
			maxLen = currLen
		}

	}

	return maxLen
}
~~~
