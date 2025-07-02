package main

import "fmt"

func main() {
	fmt.Println("=== Go运算符演示 ===")

	// 算术运算符: + - * / % ++ --
	fmt.Println("\n【算术运算符】")
	a, b := 10, 3
	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("a + b = %d\n", a+b)  // 加法: 13
	fmt.Printf("a - b = %d\n", a-b)  // 减法: 7
	fmt.Printf("a * b = %d\n", a*b)  // 乘法: 30
	fmt.Printf("a / b = %d\n", a/b)  // 除法: 3 (整数除法)
	fmt.Printf("a %% b = %d\n", a%b) // 取模: 1

	// 自增自减运算符
	a++
	fmt.Printf("a++ 后 a = %d\n", a) // 11
	b--
	fmt.Printf("b-- 后 b = %d\n", b) // 2

	// 逻辑运算符: && || !
	fmt.Println("\n【逻辑运算符】")
	x, y := true, false
	fmt.Printf("x = %t, y = %t\n", x, y)
	fmt.Printf("x && y = %t\n", x && y) // 逻辑与: false
	fmt.Printf("y || x = %t\n", y || x) // 逻辑或: true
	fmt.Printf("!x = %t\n", !x)         // 逻辑非: false
	fmt.Printf("!y = %t\n", !y)         // 逻辑非: true

	// 关系运算符: == != > < >= <=
	fmt.Println("\n【关系运算符】")
	m, n := 5, 8
	fmt.Printf("m = %d, n = %d\n", m, n)
	fmt.Printf("m == n: %t\n", m == n) // 等于: false
	fmt.Printf("m != n: %t\n", m != n) // 不等于: true
	fmt.Printf("m > n: %t\n", m > n)   // 大于: false
	fmt.Printf("m < n: %t\n", m < n)   // 小于: true
	fmt.Printf("m >= n: %t\n", m >= n) // 大于等于: false
	fmt.Printf("m <= n: %t\n", m <= n) // 小于等于: true

	// 位运算符: & | ^ << >>
	fmt.Println("\n【位运算符】")
	p, q := 12, 10 // 12=1100, 10=1010 (二进制)
	fmt.Printf("p = %d (%04b), q = %d (%04b)\n", p, p, q, q)
	fmt.Printf("p & q = %d (%04b)\n", p&q, p&q)    // 按位与: 8 (1000)
	fmt.Printf("p | q = %d (%04b)\n", p|q, p|q)    // 按位或: 14 (1110)
	fmt.Printf("p ^ q = %d (%04b)\n", p^q, p^q)    // 按位异或: 6 (0110)
	fmt.Printf("p << 2 = %d (%06b)\n", p<<2, p<<2) // 左移: 48 (110000)
	fmt.Printf("p >> 2 = %d (%02b)\n", p>>2, p>>2) // 右移: 3 (11)

	// 赋值运算符: = += -= *= /= %= &= |= ^=
	fmt.Println("\n【赋值运算符】")
	num := 10
	fmt.Printf("初始值 num = %d\n", num)

	num += 5                            // num = num + 5
	fmt.Printf("num += 5 后: %d\n", num) // 15

	num -= 3                            // num = num - 3
	fmt.Printf("num -= 3 后: %d\n", num) // 12

	num *= 2                            // num = num * 2
	fmt.Printf("num *= 2 后: %d\n", num) // 24

	num /= 4                            // num = num / 4
	fmt.Printf("num /= 4 后: %d\n", num) // 6

	num %= 4                             // num = num % 4
	fmt.Printf("num %%= 4 后: %d\n", num) // 2

	// 位赋值运算符
	bit := 12 // 1100
	fmt.Printf("\n位赋值运算 - 初始值 bit = %d (%04b)\n", bit, bit)

	bit &= 10                                        // bit = bit & 10
	fmt.Printf("bit &= 10 后: %d (%04b)\n", bit, bit) // 8 (1000)

	bit |= 5                                        // bit = bit | 5
	fmt.Printf("bit |= 5 后: %d (%04b)\n", bit, bit) // 13 (1101)

	bit ^= 3                                        // bit = bit ^ 3
	fmt.Printf("bit ^= 3 后: %d (%04b)\n", bit, bit) // 14 (1110)

	// Go没有三元运算符，但可以用if-else实现
	fmt.Println("\n【三元运算符的替代写法】")
	score := 85
	var result string
	if score >= 60 {
		result = "及格"
	} else {
		result = "不及格"
	}
	fmt.Printf("分数 %d: %s\n", score, result)

	// 多重赋值 (逗号运算符的使用)
	fmt.Println("\n【多重赋值】")
	var1, var2 := 100, 200
	fmt.Printf("var1 = %d, var2 = %d\n", var1, var2)

	// 交换变量值
	var1, var2 = var2, var1
	fmt.Printf("交换后: var1 = %d, var2 = %d\n", var1, var2)

	// 指针运算符: & *
	fmt.Println("\n【指针运算符】")
	value := 42
	ptr := &value // 获取value的地址
	fmt.Printf("value = %d\n", value)
	fmt.Printf("value的地址 &value = %p\n", &value)
	fmt.Printf("指针 ptr = %p\n", ptr)
	fmt.Printf("指针指向的值 *ptr = %d\n", *ptr)

	// 通过指针修改值
	*ptr = 100
	fmt.Printf("通过指针修改后 value = %d\n", value)

	// 运算符优先级演示
	fmt.Println("\n【运算符优先级】")
	result1 := 2 + 3*4         // 乘法优先级高于加法
	result2 := (2 + 3) * 4     // 括号改变优先级
	result3 := 10 > 5 && 3 < 7 // 关系运算符优先级高于逻辑运算符

	fmt.Printf("2 + 3 * 4 = %d\n", result1)       // 14
	fmt.Printf("(2 + 3) * 4 = %d\n", result2)     // 20
	fmt.Printf("10 > 5 && 3 < 7 = %t\n", result3) // true

	// 类型转换与运算
	fmt.Println("\n【类型转换与运算】")
	intVal := 10
	floatVal := 3.14
	// Go需要显式类型转换
	result4 := float64(intVal) + floatVal
	result5 := intVal + int(floatVal)

	fmt.Printf("int(%d) + float64(%.2f) = %.2f\n", intVal, floatVal, result4)
	fmt.Printf("int(%d) + int(%.2f) = %d\n", intVal, floatVal, result5)
}
