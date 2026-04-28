package main

import "fmt"

type Person struct {
	name, city string
	age        int8
}

func main() {
	// 结构体
	var p Person
	p.age = 18
	p.name = "张三"
	p.city = "北京"

	fmt.Printf("type: %T , value: %#v\n", p, p)

	p.age = 20
	fmt.Printf("type: %T , value: %#v\n", p, p)

	// 匿名结构体
	var user struct {
		name, city string
		age        int8
	}
	user.age = 18
	user.name = "李四"
	user.city = "上海"

	fmt.Printf("type: %T , value: %#v\n", user, user)

	// 结构体指针
	var p2 = new(Person)
	fmt.Printf("%T , %#v\n", p2, p2)

	// 通过指针访问结构体成员
	(*p2).age = 18
	(*p2).name = "王五"
	(*p2).city = "广州"

	// Go语言提供了语法糖，可以直接通过指针访问结构体成员
	p2.age = 18
	p2.name = "王五"
	p2.city = "广州"

	fmt.Printf("%T , %#v\n", p2, p2)

	// 取结构体指针的地址进行实例化
	var p3 = &Person{
		name: "赵六",
		city: "深圳",
		age:  18,
	}
	p3.age = 98
	fmt.Printf("%T , %#v\n", p3, p3)

	var p4 = Person{
		name: "卓七",
		city: "上海",
		age:  18,
	}
	p4.age = 100
	fmt.Printf("%T , %#v\n", p4, p4)

	// 结构体构造函数
	p5 := NewPerson("钱八", "杭州", 18)
	fmt.Printf("%T , %#v\n", p5, p5)
	p5.Dream()
	p5.SetCity("成都")
	fmt.Printf("%T , %#v\n", p5, p5)
}

func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言\n", p.name)
}

func (p *Person) SetCity(city string) {
	p.city = city
}

func NewPerson(name, city string, age int8) *Person {
	return &Person{
		name: name,
		city: city,
		age:  age,
	}
}
