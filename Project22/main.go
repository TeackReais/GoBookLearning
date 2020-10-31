package main

import (
	"fmt"
)

type Values map[string][]string

func main() {
	a := make(Values) //已经初始化不是nil
	var b Values      //未初始化未nil
	fmt.Println(a)
	fmt.Println(a == nil)
	fmt.Println(b)
	fmt.Println(b == nil)
//	b = make(Values)
//	fmt.Println(b)
//	fmt.Println(b == nil)
	a.Add("a", "1")
	a.Add("a", "2")
	b.Add("a", "1")
	fmt.Println(a)
}

func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}
