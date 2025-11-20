package main

import (
	"fmt"
	"sync"
)

func test(wg *sync.WaitGroup) {
	defer wg.Done() // 在函数结束时通知 WaitGroup
	for i := 0; i < 10; i++ {
		fmt.Println("test()", i)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)    // 告诉 WaitGroup 有 1 个 goroutine 要等待
	go test(&wg) // 启动 goroutine

	for i := 0; i < 10; i++ {
		fmt.Println("main()", i)
	}

	wg.Wait() // 等待所有 goroutine 完成
	fmt.Println("所有 goroutine 都完成了！")
}
