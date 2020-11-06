package main

import "fmt"

func main() {
	val := add(2, 3)
	fmt.Println(val)
	div, mul := divmul(10, 20)
	fmt.Println("div :",div)
	fmt.Println("mul :",mul)

	i := 10
	double(i)
	fmt.Println(i)
	doublePtr(&i)
	fmt.Println(i)

}

func doublePtr(i *int) { // i store address of i
	*i *= 2 // *i is dereference , *i store value of i *address -> gives value
}

func double(i int) {
	i *= 2
}

func add(i int, i2 int) int {// return single values
	return i + i2
}

func divmul(i int, j int) (int , int)  {// return 2 different values
	return i / j, i * j
}
