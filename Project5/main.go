package main

import (
	//	"strconv"
	"fmt"
)

func main() {
	//	var a int
	//	a,_=strconv.ParseInt("1024", 2, 32)
	//	fmt.Println(a)
	b := "hello 你world"
	fmt.Println(len(b))
	//	c := b[:7] //此处是对字节进行操作，而非字符
	fmt.Print("abcd\n")
	fmt.Print("ef")
	d:="abc"
	e:=[]byte(d)
	f:=string(e)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	e[0]='e'
	fmt.Println(d)
	fmt.Printf("%s",e)
	fmt.Println(f)
	//	fmt.Println(c)
}
