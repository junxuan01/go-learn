package main

import (
	"fmt"
	"os"
)

var (
	allStudents map[int64]*student // 全局变量，存储所有学生信息
)

type student struct {
	ID   int
	Name string
}

func showStudents() {
	// 显示所有学生信息
	if len(allStudents) == 0 {
		fmt.Println("没有学生信息")
		return
	}
	fmt.Println("学生信息如下:")
	for id, stu := range allStudents {
		fmt.Printf("ID: %d, Name: %s\n", id, stu.Name)
	}
	fmt.Println("=====================================")
}

func addStudent() {
	// 添加学生信息
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生ID: ")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名: ")
	fmt.Scanln(&name)

	if _, exists := allStudents[id]; exists {
		fmt.Printf("ID为%d的学生已存在，请重新输入\n", id)
		return
	}
	allStudents[id] = &student{
		ID:   int(id),
		Name: name,
	}
	fmt.Printf("学生ID为%d的学生已添加成功\n", id)
	fmt.Println("=====================================")

}

func deleteStudent() {
	// 删除学生信息
	var id int64
	fmt.Print("请输入要删除的学生ID: ")
	fmt.Scanln(&id)

	if _, exists := allStudents[id]; !exists {
		fmt.Printf("ID为%d的学生不存在\n", id)
		return
	}
	delete(allStudents, id)
	fmt.Printf("ID为%d的学生已删除成功\n", id)
	fmt.Println("=====================================")
}

func main() {
	allStudents = make(map[int64]*student, 48) // 初始化全局变量
	for {
		fmt.Println("===========欢迎光临学生管理系统===========")
		fmt.Println(`
			1. 查看学生 
			2. 添加学生
			3. 删除学生
			4. 退出
	`)
		fmt.Print("请输入你的选择: ")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择的操作是: %d\n", choice)

		switch choice {
		case 1:
			showStudents()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			fmt.Println("退出系统")
			os.Exit(1)
		default:
			fmt.Println("无效的选择，请重新输入")
		}
	}
}
