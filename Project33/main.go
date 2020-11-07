package main

import (
	"fmt"
)

var a int = 1

type example1 *int

type d struct{
	example1
}

type Employee struct {
	Name string
	Id   int
}

func main() {
	fmt.Println(a)
	change(a)
	fmt.Println(a)
	map2 := map[int]string{1: "11111"}
	change2(map2)
	for a, b := range map2 {
		fmt.Println(a, b)
	}
	var xu Employee = Employee{Name: "XUKAI", Id: 20}
	fmt.Println(xu)
	fmt.Println(xu.Id)
	fmt.Println(xu.Name)
	fmt.Println(&xu)
	//fmt.Println(xu.Id.Name)
	fmt.Println(&xu.Name)
	fmt.Println(&xu.Id)
}

func change(b int) {
	b = b + 1
}

func change2(map1 map[int]string) {
	if map1 == nil {
		return
	}
	map1[5] = "12345"
}
