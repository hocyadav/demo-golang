package main

import "fmt"

func main() {
	x:= 10
	funcName(x)
	fmt.Println(x)// print

	funName2(&x)// pass pointer
	fmt.Println(x)// print old value
}

func funName2(i *int) {
	(*i) = (*i)*2 //dereference it and use it
}

func funcName(x int) {
	x = x*2
}
