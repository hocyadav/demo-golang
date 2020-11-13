package main

import "fmt"
// https://www.geeksforgeeks.org/interfaces-in-golang/
func main() {
	var a interface{} = "hariom"
	myfun(a)

	var b interface{} = 123
	myfun(b)
}

func myfun(a interface{}) {
	extractedValue, ok := a.(string)//if interface contain string then it will store in extractedValue as value and ok as true, else false
	fmt.Println(extractedValue, ok)
}
