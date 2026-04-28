package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	ID   int // 学号唯一
	Name string
}
type class struct {
	ID       int // 班级唯一
	Title    string
	Students []student
}

func newStudent(id int, name string) student {
	return student{
		ID:   id,
		Name: name,
	}
}

func main() {
	c1 := class{
		ID:       1,
		Title:    "软件工程",
		Students: make([]student, 0, 20),
	}
	// 往班级里添加学生
	for i := range 10 {
		student := newStudent(i+1, fmt.Sprintf("stu%02d", i+1))
		c1.Students = append(c1.Students, student)
	}
	// fmt.Printf("%v\n", c1)
	data, err := json.Marshal(c1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("data: %s\n", data)

	// 反序列化
	jsonStr := `{"ID":1,"Title":"软件工程","Students":[{"ID":1,"Name":"junxuan"},{"ID":2,"Name":"Guoxiangwen"},{"ID":3,"Name":"学生3"},{"ID":4,"Name":"学生4"},{"ID":5,"Name":"学生5"},{"ID":6,"Name":"学生6"},{"ID":7,"Name":"学生7"},{"ID":8,"Name":"学生8"},{"ID":9,"Name":"学生9"},{"ID":10,"Name":"学生10"}]}`
	var c2 class
	err = json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("c2: %#v\n", c2)
}
