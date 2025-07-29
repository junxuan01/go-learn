package main

import "fmt"

// 闭包

var c = "44"

func f1(f func()) {
	fmt.Println("函数作为参数的示例")
}
func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

func main() {
	// f1(f2(2, 3))
}
