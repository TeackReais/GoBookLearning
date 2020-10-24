package main

import "fmt"
var a=100
func main(){
	fmt.Println(&a,a)
	a:=200
	fmt.Println(&a,a)
	function1()
	fmt.Println(&a,a)
}

func function1(){
	a:=300
	fmt.Println(&a,a)
}