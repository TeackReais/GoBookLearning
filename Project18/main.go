package main

import "fmt"

var a = 0

func main() {
	fmt.Println(a) //1
	a++
	fmt.Println(a)       //1
	defer fmt.Println(a) //1 最后为1，a数值已经确定
	a++
	fmt.Println(a)       //2
	fmt.Println(test(5)) //defer使函数在return之后运行
}

func test(b int) int {
	c := func() { //如果没有c这个变量，这个匿名函数不会运行
		b++
	}
	defer c()
	return b
}
