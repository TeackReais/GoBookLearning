package main

import "fmt"

func main() {
	fmt.Println(triple(4))
}

func triple(x int) (result int) {
	defer func() { x = x + x }()
	return 2 * x
}
