package main

import "fmt"

//type interface1 interface {
//	String() string
//}

func interfacejudge(input interface{}) {
	switch v := input.(type) {
	case int:
		{
			fmt.Printf("%v,%T", v, v)
			fmt.Println()
		}
	case string:
		{
			fmt.Printf("%v,%T", v, v)
			fmt.Println()
		}
	}
}

func main() {
	interfacejudge("asdasd")
	//	var inter interface1
	//	inter.String()
}
