package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name string
	Age  int
}

func (m *Monster) Set(name string, age int) {
	m.Age = age
	m.Name = name
}
func reflectTest01(input interface{}) {
	rType := reflect.TypeOf(input)
	rVal := reflect.ValueOf(input)
	Valinterface := rVal.Interface()
	fmt.Println(rType, rVal)
	Val := Valinterface.(int)
	fmt.Printf("%T\n", Val)
	fmt.Println(Val * 2)
	a := rType.Kind()
	fmt.Println(a)
	var m1 Monster
	m1.Set("abc", 10)
	fmt.Println(m1)
}

func main() {
	reflectTest01(123)
}
