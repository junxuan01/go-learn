package main

import "fmt"

func modify1(x int) {
	x = 10
}
func modify2(x *int) {
	*x = 10
}

func main() {
	//指针
	// 指针是一个变量，它存储了另一个变量的内存地址。
	// 在Go语言中，指针是一个重要的概念，它允许我们直接
	// 操作内存地址，从而实现更高效的内存管理和数据操作。
	// 指针的使用可以提高程序的性能，减少内存复制的开销。
	// 但是，指针的使用也需要小心，因为错误的指针操作
	// 可能导致程序崩溃或数据损坏。

	// a := 1
	// modify1(a)
	// println("a的值:", a) // 输出: a的值: 1

	// modify2(&a)
	// println("a的值:", a) // 输出: a的值: 10
	var n = 18
	var p = &n // 获取变量n的地址，p是一个指针变量 *int
	fmt.Printf("%v %p\n", n, &n)
	fmt.Printf("%v %p\n", p, &p)

	// var a *int
	// *a = 100
	// println("a的值:", *a)

	// var a = new(int)    // 使用new关键字分配内存，a是一个指向int类型的指针
	// *a = 100            // 给指针a指向的内存赋值
	// println("a的值:", *a) // 输出: a的值: 100

	// make是用来给slice、map和channel分配内存的函数，make函数返回的是对应的这三个类型本身

}
