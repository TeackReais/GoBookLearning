package main

import (
	"fmt"
	"os"
	"io"
	"io/ioutil"
)

type StructA struct {
	intA    int
	Str     string
	Address *StructA
}

func main() {
	fmt.Println(os.Args[1:])
	a := os.RemoveAll("/abc/")
	b := 3
	fmt.Printf("%T %T", a,b)
	readFile("a.data")
}

func readFile(fileAddress string) string {
	file, err := os.OpenFile(fileAddress, os.O_RDONLY, 0600)
	if(err!=io.EOF){
		fmt.Println("A")
	}
	defer file.Close()
	contentByte, err := ioutil.ReadAll(file)
	if(err==io.EOF){
		fmt.Println("B")
	}
	return string(contentByte)
}