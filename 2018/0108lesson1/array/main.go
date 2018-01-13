package main

import "fmt"

func main() {
	var arr [5]int // 声明了一个int类型的数组
	arr[0] = 666   // 数组下标是从0开始的
	fmt.Println("arr[0]=", arr[0])
	fmt.Println("arr[1]=", arr[1]) //默认返回0
	//遍历数组
	for item := range arr {
		fmt.Println("arr:", item)
	}

	//声明且初始化
	var arr1 = [2]string{"tom", "jerry"}
	for i := 0; i < 2; i++ {
		fmt.Printf("arr1[%d] = %s\n", i, arr1[i])
	}

}
