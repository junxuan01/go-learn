package main

import "fmt"

type person struct {
	name  string
	age   int
	hobby []string
}

func main() {
	var p person
	p.name = "张三"
	p.age = 25
	p.hobby = []string{"读书", "旅行", "编程"}
	fmt.Println("p", p)
	fmt.Printf("%T\n", p) // 输出结构体的默认格式
}
