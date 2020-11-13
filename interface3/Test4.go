package main

import "fmt"

type myinterface interface {
	InterfaceFun() int
}

func main() {
	var t myinterface// declare variable interface type

	fmt.Println("interface value ",t)// interface return dynamic value
	fmt.Printf("interface type %T", t)//nil coz this interface dont know who has implemented this interface
}
