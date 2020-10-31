package main

import (
	"bytes"
	"fmt"
	"io"
)

//var input =flag.Duration("input", value time.Duration, "1213")
const debug = true

type Employee struct {
	Name   string
	Old    int
	Salary int
}

func main() {
	var any interface{}
	any = true //any可以接受任何值
	fmt.Println(any)
	//	flag.Parse()
	//	flag.Duration(name string, value time.Duration, usage string)
	var buf io.Writer
	fmt.Println(buf)
	f(buf)
	var buf2 *bytes.Buffer
	buf2 = new(bytes.Buffer)
	f(buf2)
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("DONE!\n"))
		fmt.Println("out!=nil")
	} else {
		fmt.Println("out=nil")
	}
}

/*
func main(){
	var mayun Employee = Employee{Name: "马云"}
	var xukai *Employee
	var sunshaohui *Employee
	sunshaohui = new(Employee)
	fmt.Println(xukai)
	fmt.Println(sunshaohui)
	*sunshaohui = Employee{Name: "孙少晖", Old: 19, Salary: 0}
	fmt.Println(sunshaohui)
	fmt.Println(*sunshaohui)
	fmt.Printf("%p", mayun)
	*xukai = Employee{Name: "徐凯"}
	fmt.Println(xukai)
}*/

/*
func main(){
	var buf3 *bytes.Buffer //nil指针必须用另一个指针赋值或者用new初始化，用*赋值有问题
	var buf2 *bytes.Buffer
	if debug {
		buf2 = new(bytes.Buffer)
	} //buf2指针指向一个不知名该类型变量
	fmt.Println(buf3)
	fmt.Println(buf2)
	fmt.Println(buf3 == buf2) //一个nil一个有值不相同
}*/
