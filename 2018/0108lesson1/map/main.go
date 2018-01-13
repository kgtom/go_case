package main

import "fmt"

func main() {

	// 创建集合
	m := make(map[string]string)

	//新增 key-value
	m["France"] = "Paris"
	m["Italy"] = "Rome"

	//遍历
	for country := range m {
		fmt.Println("Capital of", country, "is", m[country])
	}

	//查找是否存在
	captial, ok := m["United States"]
	if ok {
		fmt.Println("Capital of United States is", captial)
	} else {
		fmt.Println("Capital of United States is not present")
	}

	//删除
	delete(m, "France")
	fmt.Println("Entry for France is deleted")

	fmt.Println("删除元素后 map")

	//遍历
	for country := range m {
		fmt.Println("Capital of", country, "is", m[country])
	}
}
