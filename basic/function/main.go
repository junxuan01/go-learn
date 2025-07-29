package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

// 闭包
func closureExample() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// 匿名函数

// 没有返回值的函数
func printHello() {
	println("Hello, World!")
}

// 没有参数，但有返回值的函数
func getGreeting() string {
	return "Hello, Go!"
}

// 多个返回值的函数
func getCoordinates() (int, int) {
	return 10, 20
}

// 可变参数的函数
func addNumbers(a string, b ...int) {
	fmt.Print(a)
	fmt.Println(b)
}

func main() {
	var sum = sum(1, 2)
	println("1 + 2 =", sum) // 输出: 1 + 2 = 3
	printHello()            // 输出: Hello, World!
	getGreeting()
	m, n := getCoordinates()
	fmt.Println(m, n)
	addNumbers("sdf", 4, 3, 2)
}

// Go语言中
