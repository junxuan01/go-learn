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
	// 字符串修改
	str5 := "Hello, World!"
	str5 = strings.ReplaceAll(str5, "World", "Go") // 替换
	fmt.Println("str5修改后的结果:", str5)               // 输出：str5修改后的

	s2 := "白萝卜"
	s3 := []rune(s2)                     // 将字符串转换为rune切片
	s3[0] = '红'                          // 修改第一个字符
	fmt.Println("s2修改后的结果:", string(s3)) // 输出：s2修改后的结果: 红萝卜

	c1 := "H"
	c2 := 'H'                                    // 字符串和字符的区别
	fmt.Println("c1的类型:", fmt.Sprintf("%T", c1)) // 输出：c1的类型: string
	fmt.Println("c2的类型:", fmt.Sprintf("%T", c2)) // 输出：c2的类型: int32
	fmt.Println("c1和c2是否相等:", c1 == string(c2))  // 输出：c1和c2是否相等: true

}
