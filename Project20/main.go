package main

import "fmt"

type Employee struct {
	Name   string
	Salary int
}

var XUKAI Employee

func main() {
	var a = 1
	var b *int = &a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)
	add1(a, 2)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*b)
}

func add1(a int, add int) {
	p := &a
	*p=*p+add
	return
}
