package main

import "fmt"

type Square struct {
	Length int
}
func (s *Square) InterfaceFun()  int {
	return s.Length * s.Length
}

type Circle struct {
	Radius int
}
func (c *Circle) InterfaceFun()  int{
	return 10
}

type Interface1 interface {
	InterfaceFun() int
}

func main() {
	square := &Square{
		Length: 11,
	}
	interface1 := Interface1(square)
	i := interface1.InterfaceFun()
	fmt.Println("area", i)

	area := square.InterfaceFun()
	fmt.Println(area)

	circle := &Circle{
		Radius: 134,
	}

	arry := []Interface1{square, circle}
	fmt.Println(circle)
	for index, value := range arry{
		fmt.Println(index)
		fmt.Println(value)
	}

}
