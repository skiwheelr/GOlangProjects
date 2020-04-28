package main

import (
	"fmt"
	_ "reflect"
)

type employee struct {
	name string
	salary int
}
func main() {
	//creates a list of employee structs
	employees := []employee{
		{
			name: "Joe",
			salary: 50,
		},
		{
			name: "Sarah",
			salary: 100,
		},
		{
			name: "Greg",
			salary: 150,
		},
	}
	fmt.Println(employees)
	fmt.Println(employees[1].salary)
	maxSal := 0
	maxEmp := "Null"
	for _,employee := range employees {
		if(employee.salary > maxSal){
			maxSal = employee.salary
			maxEmp = employee.name
		}
	}
	fmt.Print("The Highest Paid Employee is ",maxEmp," who earns ",maxSal)
//Prints: The Highest Paid Employee is Greg who earns 150


}
