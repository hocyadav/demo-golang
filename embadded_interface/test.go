package main

import "fmt"

type interface1 interface {
	bangBang()
}

type interface2 interface {
	hiBye()
}

type finalinterface interface {
	interface1
	interface2
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
	student.bangBang()

	student.hiBye()



}
