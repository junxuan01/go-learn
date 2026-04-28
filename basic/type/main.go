package main

import "fmt"

// 自定义类型
type MyInt int

// 类型别名
type MyIntAlias = int

func main() {
	var a MyInt = 3
	fmt.Printf("type: %T , value: %v\n", a, a)

	var b MyIntAlias = 5
	fmt.Printf("type: %T , value: %v\n", b, b)
}
