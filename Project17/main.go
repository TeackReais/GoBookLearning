package main

import (
	"fmt"
)

var a = 1

func func1() {
	fmt.Println(a)
	a++ //全局变量，修改有效
}

var func2 func() = func1 //func2是func()类型
var func3 = func() int {
	var b int
	b++
	return b * b
}

func func11() int {
	var int1 int
	int1++
	return int1 * int1
}
func main() {
	func10 := func3()
	func12 := func11()
	func2()
	func2()
	fmt.Println(func3())
	fmt.Println(func3())
	fmt.Println(func10)
	fmt.Println(func10)
	fmt.Println(func12)
	fmt.Println(func12)
	f := func4         //f为func() func() int类型
	f2 := func4()      //f2为func() int类型
	fmt.Println(f)     //f为一个函数
	fmt.Println(f2)    //f2为一个函数
	fmt.Println(func5) //func5为一个函数
	//以下为有效的静态变量
	fmt.Println(func5()) //func5()为一个函数的结果
	fmt.Println(func5())
	func5 = func4() //此步func4()函数运行并返回匿名函数，func4()运行时创建了变量c
	fmt.Println(func5())
	fmt.Println(func5())
	fmt.Println(f2())
	fmt.Println(f2())
}

var func5 func() int = func4() //同f2

func func4() func() int {
	var c int = 0
	return func() int {
		c++
		return c * c
	}
}
