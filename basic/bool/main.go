package main

import "fmt"

func main() {
	// 布尔类型
	var b1 bool = true
	var b2 bool = false

	fmt.Println("b1:", b1) // 输出：b1: true
	fmt.Println("b2:", b2) // 输出：b2: false

	// 布尔类型的默认值
	var b3 bool            // 默认值为false
	fmt.Println("b3:", b3) // 输出：b3: false

	// 布尔类型的运算
	fmt.Println("b1 && b2:", b1 && b2) // 输出：b1 && b2: false
	fmt.Println("b1 || b2:", b1 || b2) // 输出：b1 || b2: true
	fmt.Println("!b1:", !b1)           // 输出：!b1: false

	// fmt.Println("!b2:", int(b1))
}

// 注意：Go语言中没有隐式类型转换，布尔类型不能直接转换为整数
// 注意 布尔值无法参与算术运算，不能直接转换为整数或其他数值类型。
// 如果需要将布尔值转换为整数，可以使用条件表达式
// fmt.Println("b1 to int:", func(b bool) int { if b { return 1 } else { return 0 } }(b1)) // 输出：b1 to int: 1
// fmt.Println("b2 to int:", func(b bool) int { if b { return 1 } else { return 0 } }(b2)) // 输出：b2 to int: 0
