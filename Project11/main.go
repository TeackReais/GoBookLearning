package main

import (
	"fmt"
	//	"sort"
)

func main() {
	ages := make(map[string]int)
	//	ages2:=map[string]int 错误，等号右侧是一种类型，不是值
	var ages2 map[string]int
	//ages["alice"] = 31
	//ages["charlie"] = 34
	fmt.Println(ages["alice"])
	fmt.Println(ages["bob"])
	fmt.Println(ages2["bob"])
	fmt.Println(ages)
	fmt.Println(ages)
	fmt.Println(ages2 == nil)
}
