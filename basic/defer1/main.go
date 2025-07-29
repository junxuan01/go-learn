package main

import "fmt"

func main() {
	// 关键记忆：defer语句遇到就立即注册并求值参数，但函数体延迟到函数返回时才执行！
	fmt.Printf("main函数开始执行\n")
	x := 10
	defer fmt.Println("x的值为:", x) // 立即求值，保存x的值为10
	x = 20                        // 修改x的值
	// defer fmt.Printf("猜猜我会在什么时候执行？\n")
	// defer fmt.Printf("猜猜我会在什么时候执行1？\n")
	// defer fmt.Printf("猜猜我会在什么时候执行2？\n")

	fmt.Printf("main函数结束执行\n")
}
