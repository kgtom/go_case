/*
建议使用指针，原因有两个：能够改变参数的值，避免大对象的复制操作节省内存
*/
package main

import "fmt"

type Books struct {
	Id     int
	Title  string
	Author string
}

func main() {
	fmt.Println("------------------------值传递--------------------------------------------")
	var b Books
	b.Id = 1
	b.Title = "Go web 编程"
	b.Author = "astaxie"

	printBook(b)
	fmt.Println("b1 :Book title : ", b.Title, "autor:", b.Author)

	fmt.Println("-----------------------引用传递---------------------------------------------")
	printBook2(&b)
	fmt.Println("b2 :Book title : ", b.Title, "autor:", b.Author)
}

//值传递
func printBook(book Books) {

	book.Title = "GO 圣经"
	fmt.Println("printBook :Book title : ", book.Title, "autor:", book.Author)

}

//引用传递
func printBook2(book *Books) {

	book.Title = "GO 并发编程"
	fmt.Println("printBook2 :Book title : ", book.Title, "autor:", book.Author)
}
