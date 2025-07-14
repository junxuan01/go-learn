package main

import "fmt"

func main() {

	// 切片的定义
	// 什么是切片？
	// 切片是对数组的一个抽象，它提供了更灵活的数组操作方式。
	// 切片是一个动态大小的数组，可以根据需要增长或缩小。
	// 切片的底层是数组，但切片本身不需要指定长度。
	// 切片的长度可以在运行时动态变化。
	// 切片的定义方式有多种，最常见的方式是使用`make`函数或字面量方式。
	// 切片的零值是`nil`，表示一个空切片。
	// 切片的长度和容量可以通过`len`和`cap`函数获取。
	// 切片的元素可以是任何类型，包括基本类型、结构体、指针等。
	// 切片的元素可以	通过索引访问，索引从0开始。
	// 切片的元素可以通过`append`函数添加新元素。
	// 切片的元素可以通过`copy`函数复制到另一个切片
	// 切片的元素可以通过`range`关键字遍历。
	// 切片的元素可以通过`for`循环遍历。

	var s1 []int                      // 定义一个存放int类型的切片
	var s2 []string                   // 定义一个存放string类型的切片
	fmt.Println(s1, s2)               // 输出: [] []，表示两个空切片
	fmt.Println(s1 == nil, s2 == nil) // 输出: true true，表示两个切片都是nil

	var s3 = []int{1, 2, 3}
	var s4 = []string{"a", "b", "c"} // 使用字面量方式定义切片
	fmt.Println(s3, s4)              // 输出: [1 2 3] [a b c]

	// 修正：获取切片的长度和容量
	fmt.Printf("s3的长度: %d, 容量: %d\n", len(s3), cap(s3))
	fmt.Printf("s4的长度: %d, 容量: %d\n", len(s4), cap(s4))

	// 注意：这里s是数组，不是切片！
	var s = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 数组
	fmt.Printf("s的类型: %T, 长度: %d\n", s, len(s))

	// === 理解底层数组的概念 ===
	fmt.Println("\n=== 底层数组详解 ===")

	// 1. 创建切片时，Go会自动创建底层数组
	slice1 := []int{10, 20, 30, 40, 50}
	fmt.Printf("slice1: %v, 地址: %p\n", slice1, &slice1[0])

	// 2. 从数组创建切片 - 共享底层数组
	array := [5]int{1, 2, 3, 4, 5}
	slice2 := array[1:4] // 创建切片 [2, 3, 4]
	fmt.Printf("原数组: %v\n", array)
	fmt.Printf("切片slice2: %v, 长度: %d, 容量: %d\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("slice2底层数组起始地址: %p\n", &slice2[0])
	fmt.Printf("array[1]的地址: %p\n", &array[1]) // 地址相同！

	// 3. 修改切片会影响底层数组
	fmt.Println("\n修改切片对底层数组的影响:")
	slice2[0] = 999                      // 修改切片的第一个元素
	fmt.Printf("修改后的array: %v\n", array) // array也被修改了！
	fmt.Printf("修改后的slice2: %v\n", slice2)

	// 4. 多个切片共享同一个底层数组
	fmt.Println("\n多个切片共享底层数组:")
	slice3 := array[0:3] // [999, 2, 3]
	slice4 := array[2:5] // [3, 4, 5]
	fmt.Printf("slice3: %v (共享array[0:3])\n", slice3)
	fmt.Printf("slice4: %v (共享array[2:5])\n", slice4)

	// 修改slice3影响slice4
	slice3[2] = 888 // 修改array[2]
	fmt.Printf("修改slice3[2]后:\n")
	fmt.Printf("  array: %v\n", array)
	fmt.Printf("  slice3: %v\n", slice3)
	fmt.Printf("  slice4: %v\n", slice4) // slice4[0]也变了！

	// 5. 切片扩容时的底层数组变化 append
	slice5 := []string{"北京", "上海", "广州"}
	// slice5[3] = "深圳" // 直接访问索引会报错，因为切片长度不够
	fmt.Println("\n扩容前的slice5:", slice5)
	slice5 = append(slice5, "深圳") // 使用append添加元素
	fmt.Println("\n扩容后的slice5:", slice5)
	slice5 = append(slice5, "深圳", "杭州", "成都") // 再次扩容
	fmt.Println("\n再次扩容后的slice5:", slice5)

	// copy示例
	fmt.Println("\n=== copy函数示例 ===")
	slice6 := []int{1, 2, 3, 4}
	slice7 := slice6
	slice8 := make([]int, 4) // 创建一个新的切片，长度为4
	copy(slice8, slice6)
	slice6[0] = 100                // 修改slice6的第一个元素
	fmt.Println("slice6:", slice6) // [100 2 3 4]
	fmt.Println("slice7:", slice7) // [100 2 3 4] - slice7共享slice6的底层数组
	fmt.Println("slice8:", slice8) // [1 2 3 4]
	// slice8是slice6的副本，修改slice6不会影响slice8

}
