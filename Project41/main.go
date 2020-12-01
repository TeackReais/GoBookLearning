package main

import (
	"fmt"
	"os"
)

func main() {
	cmd := os.Args
	for _, c := range cmd {
		fmt.Println(c)
	}
	
}
