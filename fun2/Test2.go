package main

import fmt "fmt"

type Student struct {// constructor auto generate
	RollNo int
	Name string
	Address string
}

// constructor start with New
func NewStudent(rollNo int, name string, address string) *Student {
	// check null
	return &Student{RollNo: rollNo, Name: name, Address: address}
}

type Person struct {
	Name string
	Address string
}

func NewPerson(name string, address string) (*Person, error) {
	//return &Person{Name: name, Address: address}, nil //working
	return &Person{Name: name, Address: address}, fmt.Errorf("error from constructor..")//working
}

func main() {
	obj1 := NewStudent(12, "hari", "bangalore")

	fmt.Println(obj1)// normal print
	fmt.Printf("%+v \n",obj1)//print nice format

	person, err := NewPerson("hariom", "Delhi")
	if err == nil {
		fmt.Println(person)
		fmt.Printf("%+v \n", person)
	} else {
		fmt.Println(err)
	}
}
