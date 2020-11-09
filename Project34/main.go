package main

import (
	"fmt"
)

func main() {
	diag(inputDiag())
}

func diag(input []interface{}) { //不能用...intetface{}，那样空接口组也能放进空接口
	var length = len(input)
	var slice [][]interface{} = make([][]interface{}, length)
	for i := range slice {
		slice[i] = make([]interface{}, length)
	}
	for i, content := range input {
		fmt.Println(i, content)
		slice[i][i] = content
	}
	for i := 0; i < length; i++ {
		for d := 0; d < length; d++ {
			if slice[i][d] == nil {
				fmt.Print("0")
			} else {
				fmt.Print(slice[i][d])
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func inputDiag() []interface{} {
	fmt.Println("如果输入内容为exit，则输出")
	var input = make([]interface{}, 0)
	for {
		fmt.Println("请输入内容")
		var tempinput string
		fmt.Scanln(&tempinput)
		if tempinput == "exit" {
			break
		}
		input = append(input, tempinput)
	}
	return input
}
