package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobby   []string `json:"hobby"`
	Address Address  `json:"address"`
}
type Address struct {
	City     string `json:"city"`
	PostCode string `json:"postcode"`
}

func (address *Address) CallAddress() {
	fmt.Println("your address and postcode is:", address.City, address.PostCode)
}

var p1 = &Person{
	Name:  "Alice",
	Age:   30,
	Hobby: []string{"Reading", "Traveling"},
	Address: Address{
		City:     "New York",
		PostCode: "10001",
	},
}

func main() {
	fmt.Println("结构体的基本使用")
	fmt.Printf("Person p1: %+v\n", p1.Address.City)
	p1.Address.CallAddress()
	b, _ := json.Marshal(p1)
	fmt.Println("JSON representation of p1:", string(b))
}
