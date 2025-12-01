package main

import "fmt"

func main() {
	// map
	// Go语言中的map是一种内置的数据类型，用于存储键值对
	// 它类似于其他编程语言中的哈希表或字典。
	// map的键是唯一的，可以是任何可比较的类型，而值可以是任何类型。
	// map的主要特点是快速查找、插入和删除操作。
	// map的使用可以提高程序的性能，减少内存复制的开销。
	var m = make(map[string]int, 10) // 创建一个空的map，键为string类型
	m["apple"] = 5                   // 添加键值对
	m["banana"] = 3                  // 添加另一个键值对
	m["orange"] = 8                  // 再添加一个键值对
	fmt.Println("获取一个不存在的Key", m["not"])
	fmt.Println("map的值", m)      // 输出map的长度
	fmt.Printf("map的内容:%v\n", m) // 输出map的内容
	for k, v := range m {        // 遍历map
		println("键:", k, "值:", v) // 输出每个键值pair
	}
	delete(m, "banana")          // 删除键为"banana"的键值对
	fmt.Println("删除后map的内容:", m) // 输出删除后的map内容

	// 元素类型为map的切片
	var m1 = make([]map[int]string, 10)               // 创建一个空的切片，元素类型为map[int]string
	fmt.Println("m1的长度:", len(m1), "m1的容量:", cap(m1)) // 输出切片的长度
	m1[0] = make(map[int]string)                      // 初始化第一个元素为一个空map
	m1[0][100] = "Go语言"                               // 添加键值对到
	fmt.Println("m1的内容:", m1)                         // 输出第一个元素的内容

	// 值为切片类型的map
	var m2 = make(map[string][]int, 10) // 创建一个空的切片，元素类型为map[int]string
	m2["北京"] = []int{1, 2, 3}           // 添加键值对
	fmt.Println("m2的内容:", m2)           // 输出map的内容
}
