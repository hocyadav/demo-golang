package main

import "fmt"
// https://www.geeksforgeeks.org/interfaces-in-golang/
func main() {
	myfun("hariom")//all type is interface type with empty body :IMP
	myfun(123)
}

func myfun(a interface{}) {
	extractedValue, ok := a.(string)//if interface contain string then it will store in extractedValue as value and ok as true, else false
	fmt.Println(extractedValue, ok)
}
