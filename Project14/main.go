package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	var dilbert Employee
	var pointer *Employee = &dilbert
	pointer.Position = "123"
	var dilbert2 Employee
	dilbert2.Position = "123"
	fmt.Println(dilbert)
	fmt.Println(dilbert2)
	fmt.Println(dilbert == dilbert2)
}
