package main

import "fmt"

func main() {

	// 二维数组
	matrix := [3][...]string{
		{"beijing", "shanghai"},
		{"guangzhou", "shenzhen"},
		{"hangzhou", "nanjing"},
	}

	// 二维数组
	//
	//	var matrix = [3][2]string{
	//		{"beijing", "shanghai"},
	//		{"guangzhou", "shenzhen"},
	//		{"hangzhou", "nanjing"},
	//	}
	for i, row := range matrix {
		for j, col := range row {
			fmt.Printf("matrix[%d][%d]=%s\n", i, j, col)
		}
	}
}
