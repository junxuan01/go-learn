package main

import (
	"fmt"
	"net"
)

func main() {
	// TCP 服务器和客户端的代码请参考 tcp/server.go 和 tcp/client.go
	conn, err := net.Dial("tcp", ":8090")
	if err != nil {
		fmt.Printf("连接服务器失败, err:%v\n", err)
		return
	}
	conn.Write([]byte("Hello, Server!"))
	conn.Close()
}
