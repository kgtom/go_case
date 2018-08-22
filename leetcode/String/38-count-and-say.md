## 题目

报数序列是指一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：

~~~
1.     1
2.     11
3.     21
4.     1211
5.     111221
~~~

1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

给定一个正整数 n ，输出报数序列的第 n 项。

注意：整数顺序将表示为一个字符串。

示例 1:

~~~
输入: 1
输出: "1"
~~~

示例 2:
~~~
输入: 4
输出: "1211"
~~~

## 代码

**思路:** 获取n-1 前一个值，判断相邻字符是否相等。

~~~go
package main

import "fmt"

func main() {

	r := countAndSay(3)
	fmt.Println("r:", r)
}

func countAndSay(n int) string {
	var arr = []string{"1"}
	if n < 1 {
		return ""
	}
	if n == 1 {
		return arr[n-1]
	}

	for i := 1; i < n; i++ {
		retStr := ""
		//参照于前一个值，需获取n-1
		prevStr := arr[i-1]
		for j := 0; j < len(prevStr); j++ {
			//记录当前字符字符个数
			currCount := 1
			//获取当前字符值，用于拼接末尾值
			currVal := prevStr[j]

			//判断相邻数是否相等，相等则加1，例如：11:说明有 2个1；不相等的 则 1个currval
			for j+1 < len(prevStr) && prevStr[j+1] == prevStr[j] {
				currCount++
				j++
			}

			retStr = retStr + fmt.Sprint(currCount) + string([]byte{currVal})

		}
		arr = append(arr, retStr)
		fmt.Println("retStr:", retStr)
	}

	fmt.Println("arr:", arr)
	return arr[n-1]
}

~~~
