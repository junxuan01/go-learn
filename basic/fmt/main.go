package main

import "fmt"

func main() {
	// 查看类型
	var n = 10
	fmt.Printf("n的类型: %T\n", n)           // 输出：n的类型: int
	fmt.Printf("n的值: %d\n", n)            // 输出：n的值: 10 十进制
	fmt.Printf("n的二进制: %b\n", n)          // 输出：n的二进制: 1010
	fmt.Printf("n的八进制: %o\n", n)          // 输出：n的八进制: 12
	fmt.Printf("n的十六进制: %x\n", n)         // 输出：n的十六进制: a
	fmt.Printf("n的十六进制(大写): %X\n", n)     // 输出：n的十六进制(大写): A
	fmt.Printf("n的浮点数: %f\n", float64(n)) // 输出：n的浮点数: 10.000000
	fmt.Printf("%v\n", n)                 // 输出：n的值: 10

	fmt.Println("\n=== fmt.Printf格式化标志详细演示 ===")

	// 1. %v vs %t 的区别演示
	fmt.Printf("\n【%%v vs %%t 的重要区别】\n")
	isTrue := true
	isFalse := false

	fmt.Printf("布尔值用 %%v: %v, %v\n", isTrue, isFalse) // %v可以格式化任何类型
	fmt.Printf("布尔值用 %%t: %t, %t\n", isTrue, isFalse) // %t专门用于布尔值

	// 错误示例：%t只能用于布尔值
	number := 42
	fmt.Printf("数字用 %%v: %v\n", number) // 正确：%v万能格式
	// fmt.Printf("数字用 %%t: %t\n", number)  // 错误！%t只能用于布尔值

	// 2. %v 的万能特性演示
	fmt.Printf("\n【%%v 万能格式化标志】\n")
	str := "Hello"
	floatNum := 3.14
	arr := [3]int{1, 2, 3}
	slice := []string{"a", "b", "c"}

	fmt.Printf("字符串: %v\n", str)      // Hello
	fmt.Printf("整数: %v\n", number)    // 42
	fmt.Printf("浮点数: %v\n", floatNum) // 3.14
	fmt.Printf("布尔值: %v\n", isTrue)   // true
	fmt.Printf("数组: %v\n", arr)       // [1 2 3]
	fmt.Printf("切片: %v\n", slice)     // [a b c]

	// 3. 数字格式化演示
	fmt.Println("\n【数字格式化】")
	num := 255
	fmt.Printf("十进制 %%d: %d\n", num)    // 255
	fmt.Printf("二进制 %%b: %b\n", num)    // 11111111
	fmt.Printf("八进制 %%o: %o\n", num)    // 377
	fmt.Printf("十六进制小写 %%x: %x\n", num) // ff
	fmt.Printf("十六进制大写 %%X: %X\n", num) // FF

	// 4. 浮点数格式化演示
	fmt.Println("\n【浮点数格式化】")
	pi := 3.141592653589793
	bigNum := 123456.789

	fmt.Printf("默认浮点 %%f: %f\n", pi)        // 3.141593
	fmt.Printf("指定精度 %%.2f: %.2f\n", pi)    // 3.14
	fmt.Printf("科学计数法 %%e: %e\n", bigNum)   // 1.234568e+05
	fmt.Printf("科学计数法大写 %%E: %E\n", bigNum) // 1.234568E+05
	fmt.Printf("自动选择 %%g: %g\n", pi)        // 3.141592653589793
	fmt.Printf("自动选择 %%g: %g\n", bigNum)    // 123456.789

	// 5. 字符串格式化演示
	fmt.Println("\n【字符串格式化】")
	text := "Go语言"
	fmt.Printf("字符串 %%s: %s\n", text)  // Go语言
	fmt.Printf("带引号 %%q: %q\n", text)  // "Go语言"
	fmt.Printf("通用格式 %%v: %v\n", text) // Go语言

	// 6. 字符格式化演示
	fmt.Println("\n【字符格式化】")
	char := 'A'
	chineseChar := '中'
	fmt.Printf("字符 %%c: %c\n", char)               // A
	fmt.Printf("字符数值 %%d: %d\n", char)             // 65
	fmt.Printf("Unicode %%U: %U\n", char)          // U+0041
	fmt.Printf("中文字符 %%c: %c\n", chineseChar)      // 中
	fmt.Printf("中文Unicode %%U: %U\n", chineseChar) // U+4E2D

	// 7. 指针格式化演示
	fmt.Println("\n【指针格式化】")
	value := 100
	ptr := &value
	fmt.Printf("变量值: %v\n", value)    // 100
	fmt.Printf("指针地址 %%p: %p\n", ptr) // 0x...
	fmt.Printf("指针指向的值: %v\n", *ptr)  // 100

	// 8. 类型格式化演示
	fmt.Println("\n【类型格式化】")
	var intVar int = 42
	var strVar string = "hello"
	var boolVar bool = true
	var sliceVar []int = []int{1, 2, 3}

	fmt.Printf("int类型 %%T: %T\n", intVar)     // int
	fmt.Printf("string类型 %%T: %T\n", strVar)  // string
	fmt.Printf("bool类型 %%T: %T\n", boolVar)   // bool
	fmt.Printf("slice类型 %%T: %T\n", sliceVar) // []int

	// 9. 特殊格式化演示
	fmt.Println("\n【特殊格式化】")
	fmt.Printf("百分号 %%%%: %%\n") // %%

	// 10. 宽度和精度控制
	fmt.Println("\n【宽度和精度控制】")
	name := "Alice"
	score := 95.5

	fmt.Printf("右对齐: '%10s'\n", name)      // '     Alice'
	fmt.Printf("左对齐: '%-10s'\n", name)     // 'Alice     '
	fmt.Printf("补零: '%05d'\n", 42)         // '00042'
	fmt.Printf("小数精度: '%.1f'\n", score)    // '95.5'
	fmt.Printf("总宽度.精度: '%8.2f'\n", score) // '   95.50'

	// 11. 结构体格式化演示
	fmt.Println("\n【结构体格式化】")
	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "张三", Age: 25}
	fmt.Printf("结构体 %%v: %v\n", person)   // {张三 25}
	fmt.Printf("结构体 %%+v: %+v\n", person) // {Name:张三 Age:25}
	fmt.Printf("结构体 %%#v: %#v\n", person) // main.Person{Name:"张三", Age:25}

	fmt.Printf("\n=== 📚 联想记忆方法总结 ===\n")

	fmt.Printf("\n【英文单词联想记忆法】\n")
	fmt.Printf("%%d -> Decimal (十进制)\n")
	fmt.Printf("%%s -> String (字符串)\n")
	fmt.Printf("%%t -> True/false (布尔值)\n")
	fmt.Printf("%%T -> Type (类型)\n")
	fmt.Printf("%%v -> Value (值，万能格式)\n")
	fmt.Printf("%%p -> Pointer (指针地址)\n")
	fmt.Printf("%%c -> Character (字符)\n")
	fmt.Printf("%%f -> Float (浮点数)\n")
	fmt.Printf("%%b -> Binary (二进制)\n")
	fmt.Printf("%%o -> Octal (八进制)\n")
	fmt.Printf("%%x -> heXadecimal (十六进制小写)\n")
	fmt.Printf("%%X -> heXadecimal (十六进制大写)\n")
	fmt.Printf("%%q -> Quoted string (带引号的字符串)\n")
	fmt.Printf("%%e -> Exponential (科学计数法)\n")
	fmt.Printf("%%U -> Unicode (Unicode码点)\n")

	fmt.Printf("\n【形象联想记忆法】\n")
	fmt.Printf("%%v -> Victory! 万能胜利，什么都能格式化\n")
	fmt.Printf("%%t -> True/False 只有两个选择的布尔值\n")
	fmt.Printf("%%T -> Type 类型检查员，告诉你这是什么类型\n")
	fmt.Printf("%%d -> Digits 数字，十进制数字\n")
	fmt.Printf("%%s -> Sentence 句子，字符串\n")
	fmt.Printf("%%p -> Position 位置，内存地址位置\n")
	fmt.Printf("%%f -> Fraction 分数，小数点\n")
	fmt.Printf("%%c -> Character 字符，单个字符\n")

	fmt.Printf("\n【使用频率记忆法】\n")
	fmt.Printf("🌟🌟🌟 最常用：%%v (万能), %%d (整数), %%s (字符串)\n")
	fmt.Printf("🌟🌟   常用：%%t (布尔), %%T (类型), %%f (浮点)\n")
	fmt.Printf("🌟     偶尔用：%%p (指针), %%c (字符), %%x (十六进制)\n")

	fmt.Printf("\n【对比记忆法】\n")
	fmt.Printf("大小写区别：\n")
	fmt.Printf("  %%x vs %%X -> 十六进制小写 vs 大写\n")
	fmt.Printf("  %%e vs %%E -> 科学计数法小写e vs 大写E\n")
	fmt.Printf("相似功能：\n")
	fmt.Printf("  %%v vs %%s -> v是万能(任何类型), s是专门的字符串\n")
	fmt.Printf("  %%d vs %%b vs %%o vs %%x -> 都是整数，不同进制\n")
	fmt.Printf("  %%f vs %%e vs %%g -> 都是浮点数，不同表示法\n")

	fmt.Printf("\n【场景记忆法】\n")
	fmt.Printf("调试场景：%%T (看类型), %%#v (详细信息)\n")
	fmt.Printf("用户展示：%%s (字符串), %%d (数字), %%f (小数)\n")
	fmt.Printf("系统编程：%%p (指针), %%x (十六进制), %%b (二进制)\n")
	fmt.Printf("文本处理：%%c (字符), %%q (带引号), %%U (Unicode)\n")

	fmt.Printf("\n【实战记忆技巧】\n")
	fmt.Printf("1. 不确定用什么？先用 %%v 试试，它是万能的！\n")
	fmt.Printf("2. 布尔值专用 %%t，记住：t = true/false\n")
	fmt.Printf("3. 看类型用 %%T，记住：T = Type\n")
	fmt.Printf("4. 字符串用 %%s，记住：s = string\n")
	fmt.Printf("5. 整数用 %%d，记住：d = decimal\n")

	fmt.Printf("\n【错误避免提醒】\n")
	fmt.Printf("❌ 常见错误：用 %%t 格式化非布尔值\n")
	fmt.Printf("❌ 常见错误：用 %%d 格式化字符串\n")
	fmt.Printf("❌ 常见错误：用 %%s 格式化数字\n")
	fmt.Printf("✅ 万能解决：不确定就用 %%v\n")

	// 实际演示错误和正确用法
	fmt.Printf("\n【实战对比演示】\n")
	testNum := 42
	testStr := "hello"
	testBool := true

	fmt.Printf("正确用法：\n")
	fmt.Printf("  数字: %%d -> %d\n", testNum)
	fmt.Printf("  字符串: %%s -> %s\n", testStr)
	fmt.Printf("  布尔值: %%t -> %t\n", testBool)
	fmt.Printf("  万能: %%v -> %v, %v, %v\n", testNum, testStr, testBool)

	fmt.Printf("\n记住：%%v 是你的好朋友，什么都能格式化！\n")
}
