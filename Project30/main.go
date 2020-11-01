package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(input interface{}) {
	rType := reflect.TypeOf(input)
	rVal := reflect.ValueOf(input)
	Valinterface := rVal.Interface()
	fmt.Println(rType, rVal)
	Val := Valinterface.(int)
	fmt.Printf("%T\n", Val)
	fmt.Println(Val * 2)
}

func main() {
	reflectTest01(123)
}
