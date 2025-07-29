package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 计算字符串中的汉字数量
func countChineseCount(s string) int {
	var count int
	for _, v := range s {
		if unicode.Is(unicode.Han, v) { // 判断是否是汉字
			count++
		}
	}
	fmt.Println("汉字数量:", count)
	return count
}

// 函数作为参数的示例
func f1(f func() int) {
	ret := f() // 调用传入的函数
	fmt.Println("函数返回值:", ret)
}

// 函数作为返回值的示例
func f2() func() int {
	return func() int {
		return 42 // 返回一个匿名函数
	}
}

func main() {
	s2 := "how do you do"
	countChineseCount(s2) // 调用函数计算汉字数量
	slice1 := strings.Split(s2, " ")
	m1 := make(map[string]int, 10)
	for _, w := range slice1 {
		// 1. 如果原来的map中没有这个单词，则添加
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			// 2. 如果原来的map中有这个单词，则计数加1
			m1[w]++
		}
	}
	fmt.Println("单词计数:", m1)
}
