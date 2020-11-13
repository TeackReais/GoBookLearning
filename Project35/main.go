package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	b, _ := ReadFrom(os.Stdin, 20)
	fmt.Println(b)
	var a string
	fmt.Scanln(&a)
	fmt.Println([]byte(a))
}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}
