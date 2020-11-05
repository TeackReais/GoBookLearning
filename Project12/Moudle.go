package main //文件功能模块

import (
	"fmt"
	"os"
)

func content() {
	fmt.Println("[1]保存一个数据。")
	fmt.Println("[2]查询一个数据。")
	fmt.Println("[3]删除一个数据。")
	fmt.Println("[4]清空保存数据。")
	fmt.Println("[5]关闭程序。")
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

func filesError() { //其中一个文件被删除，或storage,dat文件被破坏
	fmt.Println("文件key.dat或文件storage.dat被破坏！")
	fmt.Println("请删除文件key.dat和文件storage.dat并重新启动程序。")
	fmt.Scanln()
	os.Exit(0)
}

func initialization() {
	fmt.Println("程序第一次打开，创建文件key.dat和storage.dat。")
	fmt.Println("请勿破坏任何一个文件。")
	judgeAndCreateFile("storage.dat")
	judgeAndCreateFile("key.dat")
	key["KeyMain"] = AesEncrypt(SessionId(), keyroot)
	jsonstr2 := map2Json(storage)
	writeFile("storage.dat", AesEncrypt(jsonstr2, key["KeyMain"]))
	key["MD5"] = getFileMd5("storage.dat")
	jsonstr := map2Json(key)
	writeFile("key.dat", AesEncrypt(jsonstr, keyroot))
	fmt.Println("创建完毕。")
}

func programStart() (judge bool) {
	judge = true
	keydat := readFile("key.dat")
	defer func() {
		catchpanic := recover()
		if catchpanic != nil {
			judge = false //defer在return后执行
		}
	}()
	jsonstr := AesDecrypt(keydat, keyroot) //无法处理解密失败的错误
	key = json2Map(jsonstr)
	if key["MD5"] != getFileMd5("storage.dat") {
		filesError()
	}
	storagedat := readFile("storage.dat")
	jsonstr2 := AesDecrypt(storagedat, key["KeyMain"])
	storage = json2Map(jsonstr2)
	return judge
}

func deleteDatabase() {
	for a, _ := range storage {
		delete(storage, a)
	}
}
