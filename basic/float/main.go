package main

import (
	"fmt"
)

func main() {
	f1 := 1.23456
	fmt.Printf("f1的类型: %T\n", f1) // 输出：f1的类型: float64
	f2 := float32(1.23456)
	fmt.Printf("f2的类型: %T\n", f2) // 输出：f2的类型: float32
}
