package main

import (
	"fmt"
)

func main() {
	map1 := make(map[string]string)
	map1["a"] = "123v"
	map1["as"] = "123f"
	map1["sa"] = "123a"
	map1["da"] = "123s"
	map1["c"] = ""
	a, b := map1["a"]//a为值,b为判断map中是否存在               
	c, d := map1["b"]
	e, f := map1["c"]
	fmt.Println(a, b)
	fmt.Println(c, d)
	fmt.Println(e, f)
	var int1 []string = make([]string, 10)
	int1[0] = "1"
	int1[1] = "2"
	for a, i := range map1 { //a为索引，i为内容
		fmt.Println(a, i)
		delete(map1, a)
	}
	for a, i := range int1 { //a为索引，i为内容，若只有1个变量，则只有索引
		fmt.Println(a, i)
	}
	fmt.Println(map1)
}
