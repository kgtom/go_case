package main

import (
	"fmt"
)

func main() {
	arr := [3]int{2, 2, 2}
	arr2 := arr
	arr2[0] = 8

	fmt.Println("arr:", arr)
	fmt.Printf("arr address: %p \n", &arr)
	fmt.Println("-------------")
	fmt.Println("arr2:", arr2)
	fmt.Printf("arr2 address: %p \n", &arr2)

	fmt.Println("-------------")
	var array = [3]int{0, 1, 2}
	var array2 = &array
	array2[2] = 5
	fmt.Println(array, *array2)
}
