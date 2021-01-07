package main

import "fmt"

type Interface1 interface {
	InterfaceFun() int
}

// class 1
type Square struct {
	Length int
}
// class 1 implements interface ~ java
func (s *Square) InterfaceFun()  int {
	return s.Length * s.Length
}

// class 2
type Circle struct {
	Radius int
}

func (c *Circle) InterfaceFun()  int{
	return 10
}

func main() {
	square := &Square{
		Length: 11,
	}
	interface1 := Interface1(square)
	area1 := interface1.InterfaceFun() //m2
	fmt.Println("area", area1)

	area := square.InterfaceFun()//m1
	fmt.Println(area)

	circle := &Circle{
		Radius: 134,
	}

	arry := []Interface1{square, circle} // array of interface
	fmt.Println(circle)
	for index, value := range arry{
		fmt.Println(index)
		fmt.Println(value)
	}

}
