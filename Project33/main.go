package main

import "fmt"

var a int = 1

func main() {
	fmt.Println(a)
	change(a)
	fmt.Println(a)
}

func change(b int) {
	b = b + 1
}
