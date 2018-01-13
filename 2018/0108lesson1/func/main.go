package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {

	if num1 > num2 {
		return num1
	}
	return num2
}
func main() {

	var a, b int = 100, 200
	r := max(a, b)

	fmt.Printf("最大值是 : %d\n", r)

	c, d := swap("first", "second")
	fmt.Println(c, d)
}
