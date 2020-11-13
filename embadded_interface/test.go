package main

import "fmt"
// create 2 interface , and one interface that contain these 2 interface
type interface1 interface {
	bangBang()
}

type interface2 interface {
	hiBye()
}

type finalinterface interface {
	interface1
	interface2
	//hiBye()
}

type Student struct {
	Name string
	Age int
	Address string
}

func (receiver Student) bangBang() {
	fmt.Println(receiver.Name)
	fmt.Println(receiver.Age)
	fmt.Println(receiver.Address)
}

func (receiver Student) hiBye() {
	fmt.Println("only print name ", receiver.Name)
}

func main() {
	student := Student{
		Name:    "hariom",
		Age:     31,
		Address: "bangalore",
	}
	// calling via object
	student.bangBang()

	student.hiBye()

	var f finalinterface
	f = &Student{
		Name:    "chandan",
		Age:     24,
		Address: "bangalore",
	}
	//calling via interface
	f.bangBang()
	f.hiBye()

	f2 := &Student{
		Name:    "chandan",
		Age:     24,
		Address: "bangalore",
	}

	f2.bangBang()
	f2.hiBye()
}
