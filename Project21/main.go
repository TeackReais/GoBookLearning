package main

import "fmt"

type Employee struct {
	Name   string
	Salary int
}

func main() {
	var XUKAI Employee
	a := &XUKAI
	XUKAI.Salary = 100
	fmt.Println(XUKAI.Salary)
	a.addSalary(200)
	fmt.Println(XUKAI.Salary)
	a.addSalary1(200)
	fmt.Println(XUKAI.Salary)
	XUKAI.addSalary(200)
	fmt.Println(XUKAI.Salary)
	a.addSalary1(200)
	fmt.Println(XUKAI.Salary)
	b:=100
	c:=&b
	fmt.Println(add2(*c))
}

func (p *Employee) addSalary(Add int) {
	p.Salary += Add
}

func (employee1 Employee) addSalary1(Add int) {
	employee1.Salary += Add
}

func add2(add int) int {
	return add + 3
}
