package main

import "fmt"

// go语言中的函数和方法

// 函数
func add(a, b int) int {
	return a + b
}

// 方法
type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func (p *Person) SetAge(age int) {
	p.Age = age
}
func SetAge2(p *Person, age int) {
	p.Age = age
}

func main() {
	fmt.Print("函数和方法的示例\n")
	var res = add(2, 3)
	fmt.Println("2 + 3 =", res)
	p := Person{Name: "Alice", Age: 30}
	p.Greet()
	p.SetAge(31)
	fmt.Println("After setting age, new age is:", p.Age)
	SetAge2(&p, 32)
	fmt.Println("After SetAge2, new age is:", p.Age)
}
