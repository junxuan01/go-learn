package main

import "fmt"

// 常量
const pi = 3.14

// 常量组
const (
	STATUS_OK    = 200
	NOT_FOUND    = 404
	STATUS_ERROR = 500
)

// 常量组中的隐式赋值
// 如果常量组中的第一个常量没有显式赋值，则后续常量会自动递增
// 如果第一个常量有显式赋值，则后续常量会与第一个常量相同
// 例如下面的 n1, n2, n3 都会被赋值为 100
const (
	n1 = 100
	n2
	n3
)
const (
	a1 = iota
	a2
	a3
)
const (
	b1 = iota
	b2 = iota
	_  = iota // 忽略这个值
	b3 = iota
)

const (
	c1 = iota
	c2 = 55
	c3 = iota
	c4 = 66
	c5
)

// iota的应用
const (
	_  = iota
	KB = 1 << (10 * iota) // 1 << (10 * 0) = 1
	MB = 1 << (10 * iota) // 1 << (10 * 1) = 1024
	GB = 1 << (10 * iota) // 1 << (10 * 2) = 1048576
	TB = 1 << (10 * iota) // 1 << (10 * 3) = 1073741824
	PB = 1 << (10 * iota) // 1 << (10 * 4) = 1099511627776
)

func main() {
	fmt.Println("pi:", pi)
	fmt.Println("STATUS_OK:", STATUS_OK)
	fmt.Println("NOT_FOUND:", NOT_FOUND)
	fmt.Println("STATUS_ERROR:", STATUS_ERROR)
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2) // n2 = 100
	fmt.Println("n3:", n3) // n3 = 100
	fmt.Println("a1:", a1) // a1 = 0
	fmt.Println("a2:", a2) // a2 = 1
	fmt.Println("a3:", a3) // a3 = 2
	fmt.Println("b1:", b1) // b1 = 0
	fmt.Println("b2:", b2) // b2 = 1
	fmt.Println("b3:", b3) // b3 = 3
	fmt.Println("c1:", c1) // c1 = 0
	fmt.Println("c2:", c2) // c2 = 55
	fmt.Println("c3:", c3) // c3 = 2
	fmt.Println("c4:", c4) // c4 = 3
	fmt.Println("c5:", c5) // c5 = 66
	fmt.Println("KB:", KB) // KB = 1024
	fmt.Println("MB:", MB) // MB = 1048576
	fmt.Println("GB:", GB) // GB = 1073741824
	fmt.Println("TB:", TB) // TB = 1099511627776
	fmt.Println("PB:", PB) // PB = 1125899906842624
}
