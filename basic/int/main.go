package main

import "fmt"

func main() {
	// 十进制
	var i1 = 101
	fmt.Printf("101的十进制: %d\n", i1)
	fmt.Printf("%b\n", i1) // 输出：1100101 二进制
	fmt.Printf("%o\n", i1) // 输出：145 八进制
	fmt.Printf("%x\n", i1) // 输出：65 十六进制

	// 八进制
	var i2 = 077           // 注意：在Go语言中，八进制以0开头
	fmt.Printf("%d\n", i2) // 输出：65
	// 十六进制
	i3 := 0x1A             // 注意：在Go语言中，十六进制以0x开头
	fmt.Printf("%d\n", i3) // 输出：26

	i4 := int8(9)
	fmt.Printf("i4的类型: %T\n", i4) // 输出：i4的类型: int8

}
