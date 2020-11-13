package main

import "fmt"

type myinterface interface {
	InterfaceFun() int
}

type Square struct {
	Length int
}
func (s *Square) InterfaceFun() int {
	return s.Length * s.Length
	//panic("implement me")
}

func main() {
	var t myinterface// declare variable interface type
	t = &Square{ // store address of class that have implemented interface method
		Length: 12,
	}
	valueFromInterface := t.InterfaceFun()//call interface method
	fmt.Println(valueFromInterface)
}
