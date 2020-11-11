package main

import (
	"errors"
	"fmt"
	"net"
)

func main() {
	var a string = "123"
	fmt.Println(cap(a))
	ServerAdd := make(map[string]string)
	ServerAdd = ReadTheFile()
	conn, err := net.Dial("tcp", ServerAdd["ServerAddress"])
	//defer conn.Close()
	if err != nil {
		fmt.Printf("服务器连接失败，错误原因为%s。\n", err.Error())
		return
	} else {
		fmt.Printf("服务器连接成功，服务器IP为%s。\n", conn.RemoteAddr().String())
	}
	for {
		fmt.Println("请输入传向服务器的数据。")
		var input string
		fmt.Scanln(&input)
		n, err := conn.Write([]byte(input))
		if err != nil {
			fmt.Printf("数据发送错误，错误原因为%s。", err.Error())
			fmt.Println("系统退出。")
			return
		} else {
			fmt.Printf("向服务器发送了%d个字节。\n", n)
			err := errors.New("无错误")
			fmt.Println(err)
		}
	}
}

func ReadTheFile() map[string]string {
	result := make(map[string]string)
	if judgeAndCreateFile("ServerAddress.json") {
		filestr := readFile("ServerAddress.json")
		result = json2Map(filestr)
		return result
	}
	result["ServerAddress"] = "47.100.91.201:20000"
	writeFile("ServerAddress.json", map2Json(result))
	return result
}
