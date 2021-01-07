package main

import "fmt"

// struct 2 types
type Person struct {// m1
	Name string
	Age string
}

type Employee struct { // m2
	Name, EmpId string
}

//methods 2 types
func (p Person) foo() string {// value receiver
	return "hello "+ p.Name
}

func (e *Employee) foo2() {// pointer receiver
	e.Name = "new name"
}

func main() {
	person := Person{
		Name: "hariom",
		Age:  "30",
	}

	fmt.Println(person)
	fmt.Println(person.foo())

	employee := Employee{
		Name:  "Haroim yadav",
		EmpId: "1234",
	}

	fmt.Println(employee)
	employee.foo2()
	fmt.Println(employee)
}
