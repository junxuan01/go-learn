package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string = "1000"
	b, err := strconv.ParseInt(a, 10, 64)
	if err != nil {
		fmt.Println("strconv.ParseInt err=", err)
		return
	}
	fmt.Printf("b=%v\n", b)
	boolStr := "true"
	boolValue, err := strconv.ParseBool(boolStr)
	if err != nil {
		fmt.Println("strconv.ParseBool err=", err)
		return
	}
	fmt.Printf("boolValue=%v\n", boolValue)

	// fileObj, err := os.Open("./../../basic/file/main.go")
	// if err != nil {
	// 	fmt.Println("os.Open err=", err)
	// 	return
	// }
	// fmt.Printf("fileObj=%v\n", fileObj)
	// defer fileObj.Close()
}
