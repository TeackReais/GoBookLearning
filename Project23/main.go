package main

import (
	"fmt"
)

type Phone interface {
	call()
	call2()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}
func (nokiaPhone *NokiaPhone) call2() {
	fmt.Println("I am Nokia*, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}
func (iPhone *IPhone) call2() {
	fmt.Println("I am iPhone*, I can call you!")
}
func main() {
	var phone1 Phone
	phone1 = new(NokiaPhone)
	phone1.call()
	phone1 = new(IPhone)
	phone1.call()
//	var iphone1 IPhone={}
	var iphonep *IPhone=new(IPhone)
	phone1 = iphonep
	phone1.call()
	phone1.call2()
	//	var a IPhone
	//	phone = a 如果修改第30行定义为IPhone的指针，则该语句编译错误
	//	phone = new(Mi)
	//	phone.call()
	//	phone = new(int)
	//	phone.call()
	//	var a Mi = "123"
	//	a.call()
}

/*type Mi string

func (a Mi) call() {
	fmt.Printf("I am %s, I can call you!", a)
}
func (a int) call() {
	fmt.Printfln("I am 123, I can call you!")
}*/
