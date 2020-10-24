package main

import (
	"fmt"
	"math"
)

func main() {
	//	s:="hello,world"
	//	t := "hhllo，world"
	//	var tp = &t
	//	fmt.Println(len(s))
	//	fmt.Println(len(t))
	//	fmt.Println(len(s))
	//	a:=123
	b := math.Sqrt(-1)
	//	fmt.Println(t)
	//	fmt.Println(tp)
	//	fmt.Println(*tp)
	//	fmt.Println("%d[1]%[1]f",float64(a))
	fmt.Print("%v", b) //2个元素分别打印，不带换行符
	//	fmt.Println("%v", b) //2个元素分别打印，带有换行符
	//	fmt.Printf("%v", b) //不带换行符
	defer fmt.Println(function1(1, 2))
	//只有printf才能使用以%开头的转义符
}

func function1(a int64, b int64) (bool, bool) {
	return true, true
}
