package main

import (
	"fmt"
)

func main() {
	strVar := "hello..hari"
	fmt.Println(strVar)
	//int , 64, 32, unsigned
	intAddition()
	changeIntTofloatAddition()
	floatRecommandedWay()
	multipleDeclarationInSingleLine()

	a := 123
	s := fmt.Sprintf("%d",a)// convert int to str: Sprintf return string
	fmt.Printf("value : %s Type : %T \n", s, s)// value of str access via %s, value of int access via %d
	fmt.Printf("value : %q Type : %T \n", s, s)// value of str access via %s, and %q add double quote both side

}

func multipleDeclarationInSingleLine() {
	x, y := 1.0, 2.0
	fmt.Printf("value of x : %v\n", x)
	fmt.Printf("type of x : %T\n", x)
	fmt.Printf("value of x %v, type of x : %T\n", x, x)
	mean := (x + y) / 2
	fmt.Printf("result :  %v, type of mean : %T\n", mean, mean)
}

func floatRecommandedWay() {
	x := 1.0
	y := 2.0

	fmt.Printf("value of x : %v\n", x)
	fmt.Printf("type of x : %T\n", x)
	fmt.Printf("value of x %v, type of x : %T\n", x, x)

	mean := (x + y) / 2
	fmt.Printf("result :  %v, type of mean : %T\n", mean, mean)
}

func changeIntTofloatAddition() {
	var x float64
	var y float64
	x = 1
	y = 2

	fmt.Printf("value of x : %v\n", x)
	fmt.Printf("type of x : %T\n", x)
	fmt.Printf("value of x %v, type of x : %T\n", x, x)

	var mean float64
	mean = (x + y) / 2
	fmt.Printf("result :  %v, type of mean : %T\n", mean, mean)
}

func intAddition() {
	var x int
	var y int
	x = 1
	y = 2
	fmt.Println(x)
	fmt.Println(y)

	fmt.Printf("value of x : %v", x)
	fmt.Println()
	fmt.Printf("type of x : %T", x)
	fmt.Println()
	fmt.Printf("value of x %v, type of x : %T", x, x)

	fmt.Println()

	var mean int
	mean = (x + y) / 2
	fmt.Printf("result :  %v, type of mean : %T", mean, mean)
}
