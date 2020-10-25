package main //练习4.7

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scanf("%s", &input)
	var bytes []byte = []byte(input)
	fmt.Printf("%s", reverse(bytes))
}

func reverse(inputbyte []byte) []byte {
	lens := len(inputbyte)
	for i := 0; i < lens-1; {
		inputbyte[i], inputbyte[lens-1] = inputbyte[lens-1], inputbyte[i]
		i = i + 1
		lens = lens - 1
	}
	return inputbyte
}
