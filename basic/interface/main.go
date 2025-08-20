package main

import "fmt"

type cat struct {
}
type dog struct {
}

func (c cat) Speak() {
	fmt.Print("喵喵喵\n")
}
func (d dog) Speak() {
	fmt.Print("汪汪汪\n")
}

type Animal interface {
	Speak()
}

func MakeSound(a Animal) {
	a.Speak()
}

func main() {

}
