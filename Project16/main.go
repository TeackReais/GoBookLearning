package main

import (
	"fmt"
)

func main() {
	f := square
	fmt.Printf("%T\n", f)
	fmt.Println(f(2))
	var f2 func(n int) int
	fmt.Printf("%T\n", f2)
	fmt.Printf("%v\n", f2)
	f2 = square
	fmt.Printf("%T\n", f2)
	fmt.Printf("%v\n", f2)
	fmt.Println(f2==f)
	fmt.Println(f2(2))
}
func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }
