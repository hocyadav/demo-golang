package main

import "fmt"
// https://www.geeksforgeeks.org/interfaces-in-golang/
func main() {
	var a interface{} = "hariom"
	myfun(a)
}

func myfun(a interface{}) {
	extractedValue := a.(string)
	fmt.Println(extractedValue)
}
