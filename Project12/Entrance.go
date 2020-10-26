package main //文件入口

import (
	"fmt"
	"os"
)

const keyroot = "hT7+5d0gf~&vz2OxUc#@!xtJr6^$!x" //24个字符不报错

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
		jsonstr2 := map2Json(storage)
		writeFile("storage.dat", AesEncrypt(jsonstr2, key["KeyMain"]))
		key["MD5"] = getFileMd5("storage.dat")
		jsonstr := map2Json(key)
		writeFile("key.dat", AesEncrypt(jsonstr, keyroot))
		fmt.Println("创建完毕。")
	}
	for {
		var input string
		content()
		fmt.Print("请输入命令:")
		fmt.Scanln(&input)
		if input == "4" {
			break
		} else if input == "3" {
			var keystring = inputkeystring()
			if storage[keystring] == "" {
				fmt.Println("无对应数据。")
				pauseAndReturn()
			} else {
				delete(storage, keystring)
				updateTwoFiles(storage, key)
				fmt.Println("删除完毕。")
				pauseAndReturn()
			}
			continue
		} else if input == "2" {
			var keystring = inputkeystring()
			if storage[keystring] == "" {
				fmt.Println("无数据。")
				pauseAndReturn()
			} else {
				fmt.Printf(keystring)
				println("的对应数据为:")
				fmt.Println(storage[keystring])
				pauseAndReturn()
			}
			continue
		} else if input == "1" {
			var keystring = inputkeystring()
			var inputstorage1 = inputstorage()
			if storage[keystring] != "" {
				fmt.Println("此关键字符已经有对应数据，确认要覆盖吗？")
				fmt.Println("输入Y或y覆盖，输入其它字符取消")
				var inputConfirm string
				fmt.Scanln(&inputConfirm)
				if (inputConfirm != "Y") && (inputConfirm != "y") {
					fmt.Println("已取消。")
					pauseAndReturn()
					continue
				}
			}
			storage[keystring] = inputstorage1
			updateTwoFiles(storage, key)
			fmt.Println("保存完毕。")
			pauseAndReturn()
			continue
		} else {
			fmt.Println("输入数据无效，请重新输入。")
			pauseAndReturn()
			continue
		}
	}
}
