package main

import "fmt"

type person struct {
	name  string
	age   int
	hobby []string
}

func f2(x *person) {
	x.name = "李四"
}

func main() {
	// var p person
	// var p1 = new(person)
	// p1.name = "王五"
	// fmt.Printf("%p\n", p1)
	// p.name = "张三"
	// p.age = 25
	// p.hobby = []string{"读书", "旅行", "编程"}
	// fmt.Println("p", p)
	// f2(&p)
	// fmt.Println("修改后的p", p)

	// 定义一个int类型的变量a，并输出其地址
	var a int = 10
	var b = &a
	fmt.Printf("%p\n", &a)
	fmt.Printf("%v\n", *b)
}
