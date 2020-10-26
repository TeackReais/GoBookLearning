package main

import(
	"fmt"
)

func main(){
	map1:=make(map[string]string)
	map1["a"]="123v"
	map1["as"]="123f"
	map1["sa"]="123a"
	map1["da"]="123s"
	for a,i:=range map1{
		fmt.Println(a,i)
		delete(map1, a)
	}
	fmt.Println(map1)
}