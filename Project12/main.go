package main //使用map的相关知识制作一个关键词查询器

import (
	"fmt"
	"io/ioutil"
	"os"
)

const keyroot = "hT7+5d0gf~v2OUcxtJr6^$!x" //24个字符不报错

var storage map[string]string = make(map[string]string)
var key map[string]string = make(map[string]string)

func main() {
	storagebool := judgeFileExist("storage.dat")
	keybool := judgeFileExist("key.dat")
	if storagebool && keybool {
		keydat := readFile("key.dat")
		jsonstr := AesDecrypt(keydat, keyroot) //无法处理解密失败的错误
		key = json2Map(jsonstr)
		if key["MD5"] != getFileMd5("storage.dat") {
			fmt.Println("文件key.dat或文件storage.dat被破坏！")
			fmt.Println("请删除文件key.dat和文件storage.dat并重新启动程序。")
			fmt.Scanln()
			os.Exit(0)
		}
		storagedat := readFile("storage.dat")
		jsonstr2 := AesDecrypt(storagedat, key["KeyMain"])
		storage = json2Map(jsonstr2)
	} else if storagebool || keybool {
		fmt.Println("文件key.dat或文件storage.dat被破坏！")
		fmt.Println("请删除文件key.dat和文件storage.dat并重新启动程序。")
		fmt.Scanln()
		os.Exit(0)
	} else {
		fmt.Println("程序第一次打开，创建文件key.dat和storage.dat。")
		fmt.Println("请勿破坏任何一个文件。")
		judgeAndCreateFile("storage.dat")
		judgeAndCreateFile("key.dat")
		key["KeyMain"] = AesEncrypt(SessionId(), keyroot)
		key["MD5"] = getFileMd5("storage.dat")
		jsonstr := map2Json(key)
		writeFile("key.dat", AesEncrypt(jsonstr, keyroot))
		fmt.Println("创建完毕。")
	}
	for {
		var input string
		fmt.Println("[1]保存一个数据。")
		fmt.Println("[2]查询一个数据。")
		fmt.Println("[3]删除一个数据。")
		fmt.Println("[4]关闭程序。")
		fmt.Print("请输入命令:")
		fmt.Scanln(&input)
		if input == "4" {
			break
		} else if input == "3" {
			var keystring = inputkeystring()
			if storage[keystring] == "" {
				fmt.Println("无对应数据。")
				fmt.Scanln()
			} else {
				delete(storage, keystring)
				jsonstr := map2Json(storage)
				writeFile("storage.dat", AesEncrypt(jsonstr, key["KeyMain"]))
				key["MD5"] = getFileMd5("storage.dat")
				writeFile("key.dat", AesEncrypt(map2Json(key), keyroot))
				fmt.Println("删除完毕。")
				fmt.Scanln()
			}
			continue
		} else if input == "2" {
			var keystring = inputkeystring()
			if storage[keystring] == "" {
				fmt.Println("无数据")
				fmt.Scanln() //停顿
			} else {
				fmt.Println(storage[keystring])
				fmt.Scanln()
			}
			continue
		} else if input == "1" {
			var keystring = inputkeystring()
			var inputstorage1 = inputstorage()
			if storage[keystring] != "" {
				fmt.Println("此关键字符已经有对应数据，确认要覆盖吗？")
				fmt.Println("输入Y覆盖，输入其它字符取消")
				var inputConfirm string
				fmt.Scanln(&inputConfirm)
				if inputConfirm == "Y" {
					storage[keystring] = inputstorage1
					jsonstr := map2Json(storage)
					writeFile("storage.dat", AesEncrypt(jsonstr, key["KeyMain"]))
					key["MD5"] = getFileMd5("storage.dat")
					writeFile("key.dat", AesEncrypt(map2Json(key), keyroot))
					fmt.Println("保存完毕。")
					fmt.Scanln()
				} else {
					fmt.Println("已取消。")
				}
			} else {
				storage[keystring] = inputstorage1
				jsonstr := map2Json(storage)
				writeFile("storage.dat", AesEncrypt(jsonstr, key["KeyMain"]))
				key["MD5"] = getFileMd5("storage.dat")
				writeFile("key.dat", AesEncrypt(map2Json(key), keyroot))
				fmt.Println("保存完毕。")
				fmt.Scanln()
			}
			continue
		} else {
			fmt.Println("输入数据无效，请重新输入。")
			continue
		}
	}
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

func readFile(fileAddress string) string {
	file, err := os.OpenFile(fileAddress, os.O_RDONLY, 0600)
	defer file.Close()
	contentByte, err := ioutil.ReadAll(file)
	errorExit(err, 2)
	return string(contentByte)
}

func judgeAndCreateFile(fileAddress string) bool {
	if judgeFileExist(fileAddress) {
		return true
	} else {
		file, err := os.Create(fileAddress)
		defer file.Close()
		errorExit(err, 1)
		return false
	}
}

func judgeFileExist(fileAddress string) bool {
	file, err := os.OpenFile(fileAddress, os.O_RDONLY, 0600)
	defer file.Close()
	if err != nil {
		return false
	} else {
		return true
	}
}
func writeFile(fileAddress string, context string) {
	file, err := os.OpenFile(fileAddress, os.O_WRONLY|os.O_TRUNC, 0600)
	defer file.Close()
	errorExit(err, 3)
	_, err = file.Write([]byte(context))
	errorExit(err, 4)
}

func errorExit(err error, returnnumber int) {
	if err != nil {
		fmt.Println(err.Error())
		fmt.Scanln()
		fmt.Println("按任意键退出程序。")
		os.Exit(returnnumber)
	}
}
