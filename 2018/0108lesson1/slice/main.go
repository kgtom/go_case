package main

import "fmt"

func main() {
	var aa = []int{9, 8}
	var bb = make([]int, 2)
	bb[0] = 7
	bb[1] = 6
	printSlice(aa)
	printSlice(aa)
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

	//append
	// greeting := make([]string, 3, 5)
	// // 3 is length - number of elements referred to by the slice
	// // 5 is capacity - number of elements in the underlying array

	// greeting[0] = "Good morning!"
	// greeting[1] = "Bonjour!"
	// greeting[2] = "buenos dias!"
	// greeting = append(greeting, "Suprabadham")
	//fmt.Println(greeting[3])

	//del
	//mySlice := []string{"Monday", "Tuesday"}
	// myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}

	// mySlice = append(mySlice, myOtherSlice...)
	// fmt.Println(mySlice)

	// mySlice = append(mySlice[:2], mySlice[3:]...)
	// fmt.Println(mySlice)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
