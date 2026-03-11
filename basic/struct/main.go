package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {
	var p = &Person{name: "张三3", age: 17}
	fmt.Printf("%p\n", p)
	var p1 = Person{name: "张三", age: 18}
	fmt.Printf("%#v\n", p1)
	var p2 = &Person{name: "李四", age: 20}
	fmt.Printf("%#v\n", p2)
}
