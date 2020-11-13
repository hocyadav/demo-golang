package main

import "fmt"

// https://www.geeksforgeeks.org/interfaces-in-golang/
func main() {
	myfun("hariom")//all type is interface type with empty body :IMP
	myfun(123)
	myfun(123.12)
}

func myfun(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("type is int ", a.(int))
	case string:
		fmt.Println("type is string")
	default :
		fmt.Println("default case neither int nor string")
	}
	
}
