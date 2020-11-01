package main

import "fmt"

func main() {
	x := 8
	funcName1(x)
	funcName2(x)
	funcName3()

	y := 3
	funcNameSwitch(y)
	switchWithoutCase(y)
	//for loop same as java
	forLoopSameAsJava()
}

func forLoopSameAsJava() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("------------")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i > 3 {
			break
		}
	}
	fmt.Println("------------")

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i < 3 {
			continue
		}
	}
	fmt.Println("------------")

	i := 0
	for i < 10 { // this is same as do while loop
		fmt.Println(i)
		i++
	}

	// this is like "while true"
	b := 0
	for {
		if b > 2 {// exit condition
			fmt.Println(2)
			break
		}
		fmt.Println(2)
		b++ // increment
	}
}

func switchWithoutCase(y int) {
	switch {
	case y == 1:
		fmt.Println("equal")
	case y > 1:
		fmt.Println("greater")
	default:
		fmt.Println("default")
	}
}

func funcNameSwitch(y int) { //break not required
	switch y {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("default")
	}
}

func funcName3() {
	a := 11.0
	b := 20.0
	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more than half of b")
	}
}

func funcName2(x int) {
	if x > 5 || x == 8 {
		fmt.Println("x is ", x)
	} else {
		fmt.Println("x is small")
	}
}

func funcName1(x int) {
	if x > 5 {
		fmt.Println("x is big")
	} else {
		fmt.Println("x is small")
	}
}
