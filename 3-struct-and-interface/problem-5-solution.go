package main

import "fmt"

type Employee struct {
	name   string
	age    int
	salary float64
}

func (e *Employee) GiveRaise(percentage float64) {
	e.salary += e.salary * (percentage / 100.0)
}

func main() {
	employee := Employee{
		name:   "John Doe",
		age:    30,
		salary: 50000,
	}
	fmt.Printf("Before Raise: %s's salary is %.2f\n", employee.name, employee.salary)

	// Give a raise of 10%
	employee.GiveRaise(10)

	fmt.Printf("After Raise: %s's salary is %.2f\n", employee.name, employee.salary)

}
