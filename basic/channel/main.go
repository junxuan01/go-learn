package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 10
	ch <- 21
	ch <- 32
	fmt.Printf("len(ch): %d, cap(ch): %d\n", len(ch), cap(ch)) // 分别是当前元素数量与容量
	a := <-ch
	fmt.Printf("len(ch): %d, cap(ch): %d\n", len(ch), cap(ch))
	c := <-ch
	fmt.Println("a:", a)
	fmt.Println("c:", c)
}
