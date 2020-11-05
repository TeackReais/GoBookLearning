package main //文件入口

import (
	"fmt"
)

const keyroot = "hT7+5d0gf~&vz2OxUc#@!xtJr6^$!x" //根密钥

var storage map[string]string = make(map[string]string)
var key map[string]string = make(map[string]string)

func main() {
	storagebool := judgeFileExist("storage.dat")
	keybool := judgeFileExist("key.dat")
	if storagebool && keybool {
		if !programStart() {
			filesError()
		} //导入map
	} else if storagebool || keybool {
		filesError() //文件被破坏
	} else {
		initialization() //初始化
	}
	for {
		var input string
		content()
		fmt.Print("请输入命令:")
		fmt.Scanln(&input)
		if input == "5" {
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
				fmt.Println("输入Y或y覆盖，输入其它字符取消。")
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
		} else if input == "4" {
			fmt.Println("输入Y或y确认删除保存数据，输入其它字符取消。")
			var inputConfirm string
			fmt.Scanln(&inputConfirm)
			if (inputConfirm != "Y") && (inputConfirm != "y") {
				fmt.Println("已取消。")
				pauseAndReturn()
				continue
			}
			deleteDatabase()
			updateTwoFiles(storage, key)
			fmt.Println("删除完毕。")
			pauseAndReturn()
			continue
		} else {
			fmt.Println("输入数据无效，请重新输入。")
			pauseAndReturn()
			continue
		}
	}
}
