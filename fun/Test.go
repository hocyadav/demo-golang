package main

import "fmt"

type Order struct {
	Val1 int
	Val2 string
}
// syntax : receiver + fun name with input argument [optional]+ return type [optional]
func (o *Order) fun1() int {//always pass address
	fmt.Println(o.Val1)
	return o.Val1 + 10
}

func main() {
	order := Order{
		Val1: 123,
		Val2: "hariom",
	}
	x := order.fun1()//class with method fun1
	fmt.Println(x)
}
