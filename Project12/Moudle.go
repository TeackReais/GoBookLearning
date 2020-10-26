package main

import (
	"fmt"
)

func content() {
	fmt.Println("[1]保存一个数据。")
	fmt.Println("[2]查询一个数据。")
	fmt.Println("[3]删除一个数据。")
	fmt.Println("[4]关闭程序。")
}

func inputkeystring() string {
	var inputkeystring string
	fmt.Print("请输入关键字符:")
	fmt.Scanln(&inputkeystring)
	return inputkeystring
}

func inputstorage() string {
	var inputstorage string
	fmt.Print("请输入保存字符:")
	fmt.Scanln(&inputstorage)
	return inputstorage
}

func pauseAndReturn() {
	fmt.Println("按任意键继续。")
	fmt.Scanln()
}

func updateTwoFiles(mapstorage map[string]string, mapkey map[string]string) {
	writeFile("storage.dat", AesEncrypt(map2Json(mapstorage), mapkey["KeyMain"]))
	mapkey["MD5"] = getFileMd5("storage.dat")
	writeFile("key.dat", AesEncrypt(map2Json(key), keyroot))
}
