package main

import (
	"fmt"
	"github.com/gogf/gf/crypto/gaes"
)

const keyroot = "nS8$t-5E!xueO270lIZ#4d+Uczs612"

func main() {
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
	cryptoed, err := AES(input, keyroot,16)
	fmt.Println(cryptoed)
	fmt.Println([]byte(cryptoed))
	fmt.Println(err)
	return
}

func triple(x int) (result int) {
	defer func() { x = x + x }()
	return 2 * x
}

func AES(plainText string, key string, bit int) (string, error) {
	var result []byte
	var input []byte = []byte(plainText)
	var inputkey []byte = []byte(key)
	if len(inputkey) < bit {
		for len(inputkey) < bit {
			inputkey = append(inputkey, '0')
		}
	} else if len(inputkey) > bit {
		inputkey = inputkey[:bit]
	}
	result, err := gaes.Encrypt(input, inputkey)
	if err != nil {
		return "", err
	}
	return string(result), err
}

func DEAES(cipherText string, key string, bit int) (string, error) {
	var result []byte
	var input []byte = []byte(cipherText)
	var inputkey []byte = []byte(key)
	if len(inputkey) < bit {
		for len(inputkey) < bit {
			inputkey = append(inputkey, '0')
		}
	} else if len(inputkey) > bit {
		inputkey = inputkey[:bit]
	}
	result, err := gaes.Decrypt(input, inputkey)
	if err != nil {
		return "", err
	}
	return string(result), err
}
