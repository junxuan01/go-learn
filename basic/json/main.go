package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var s = Student{Name: "张三", Age: 20}
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	fmt.Printf("json.Marshal b=%s\n", b)

}
