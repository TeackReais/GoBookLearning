package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "10.10.10.244:20000")
	defer listen.Close()
	if err != nil {
		fmt.Printf("服务器监听失败，错误为%s。\n", err.Error())
		return
	} else {
		fmt.Println("服务器正在监听", listen.Addr().String())
	}
	var clients int
	for {
		fmt.Println("正在等待下一个客户端的连接。")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("来自%s的连接失败，失败原因为%s。\n", conn.RemoteAddr().String(), err.Error())
		} else {
			fmt.Printf("服务器与%s成功建立连接。\n", conn.RemoteAddr().String())
			clients++
		}
		go process(conn, &clients)
		fmt.Printf("当前服务器连接客户端数量为%d。\n", clients)
	}
}

func process(conn net.Conn, pclients *int) {
	defer fmt.Printf("当前服务器连接客户端数量为%d。\n", *pclients)
	defer func() {
		*pclients = *pclients - 1
	}()
	defer conn.Close()
	for {
		fmt.Println("服务器正在等待来自客户端的数据。")
		receive := make([]byte, 1024)
		n, err := conn.Read(receive)
		if err != nil {
			fmt.Printf("服务器读取数据失败，错误为%s。\n", err.Error())
			break
		}
		if string(receive[:n]) == "exit" {
			break
		}
		fmt.Printf("客户端%s发送消息:\n", conn.RemoteAddr().String())
		fmt.Println(string(receive[:n]))
	}
}
