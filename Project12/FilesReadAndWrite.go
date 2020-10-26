package main //保存通用的文件读取和错误返回

import (
	"fmt"
	"io/ioutil"
	"os"
)

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
