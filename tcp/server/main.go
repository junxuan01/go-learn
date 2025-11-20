package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	// 与客户端通信
	var tmp [128]byte
	n, err := conn.Read(tmp[:])
	fmt.Printf("收到数据: %s\n", tmp[:n])
	if err != nil {
		fmt.Printf("读取数据失败, err:%v\n", err)
		return
	}
	fmt.Println("读取到数据:", string(tmp[:n]))
}

func main() {
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Printf("监听失败, err:%v\n", err)
		panic(err)
	}
	//
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("建立连接失败, err:%v\n", err)
			panic(err)
		}
		go handleConnection(conn)
	}

}
