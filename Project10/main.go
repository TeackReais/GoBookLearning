package main

import (
	"fmt"
)

func main() {
	var var1 []string=[]string{"!","2"}
	var var2 [3]string=[3]string{"!","2","123"}
	var3 := make([]string, 10, 10)
	fmt.Printf("%p\n",var1)
	fmt.Printf("%p\n",var2)
	fmt.Printf("%p\n",var3)
	var1 = var3
	fmt.Printf("%p\n",var1)
	fmt.Printf("%p\n",var3)
	a:=var2[0:3]
	b:=var2[1:3]
	c:=var2[2:3]
	fmt.Printf("%p\n",a)
	fmt.Printf("%p\n",b)
	fmt.Printf("%p\n",c)
}
