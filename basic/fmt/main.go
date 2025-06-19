package main

import "fmt"

func main() {
	// 查看类型
	var n = 10
	fmt.Printf("n的类型: %T\n", n)           // 输出：n的类型: int
	fmt.Printf("n的值: %d\n", n)            // 输出：n的值: 10 十进制
	fmt.Printf("n的二进制: %b\n", n)          // 输出：n的二进制: 1010
	fmt.Printf("n的八进制: %o\n", n)          // 输出：n的八进制: 12
	fmt.Printf("n的十六进制: %x\n", n)         // 输出：n的十六进制: a
	fmt.Printf("n的十六进制(大写): %X\n", n)     // 输出：n的十六进制(大写): A
	fmt.Printf("n的浮点数: %f\n", float64(n)) // 输出：n的浮点数: 10.000000
	fmt.Printf("%v\n", n)                 // 输出：n的值: 10

}
