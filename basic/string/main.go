package main

import (
	"fmt"
	"strings"
)

func main() {
	// 多行字符串
	var str1 = `Hello,   World!`
	fmt.Println("str1:", str1) // 输出：str1: Hello,\nWorld!
	// 字符串相关操作
	var str2 = "0123456789"
	fmt.Println("str2:", str2)                             // 输出：str2: Hello, World!
	fmt.Println("str2的长度:", len(str2))                     // 输出：str2的长度: 13
	fmt.Println("str2的第一个字符:", string(str2[0]))            // 输出：str2的第一个字符: H
	fmt.Println("str2的最后一个字符:", string(str2[len(str2)-1])) // 输出：str2的最后一个字符: !
	fmt.Println("str2的子串:", str2[7:10])                    // 输出：str2的子串: World
	fmt.Println("str2的子串(从开始到第5个字符):", str2[:5])           // 输出：str2的子串(从开始到第5个字符): Hello
	fmt.Println("str2的子串(从第7个字符到结束):", str2[7:])           // 输出：str2的子串(从第7个字符到结束): World!

	// 字符串分割
	var str4 = "Hello,World,Go"
	parts := strings.Split(str4, ",")
	fmt.Println("str4分割后的结果:", parts) // 输出：str4分割后的结果: [Hello World Go] 相当于js中的	// str4.split(",")
	//
}
