package file

import (
	"io/ioutil"
	"os"
)

//ReadFile 读取文件并返回字符串和错误
func ReadFile(fileAddress string) ([]byte, error) {
	file, err := os.OpenFile(fileAddress, os.O_RDONLY, 0600)
	defer file.Close()
	contentByte, err := ioutil.ReadAll(file)
	return contentByte, err
}

//JudgeFileExist 判断文件是否存在并返回布尔值
func JudgeFileExist(fileAddress string) bool {
	file, err := os.OpenFile(fileAddress, os.O_RDONLY, 0600)
	defer file.Close()
	if err != nil {
		return false
	}
	return true
}

//WriteFile 向文件写入字符串并返回错误
func WriteFile(fileAddress string, content []byte) error {
	file, err := os.OpenFile(fileAddress, os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		return err
	}
	return nil
}

func JudgeAndCreateFile(fileAddress string) (bool, error) {
	if JudgeFileExist(fileAddress) {
		return true, nil
	} else {
		file, err := os.Create(fileAddress)
		defer file.Close()
		return false, err
	}
}
