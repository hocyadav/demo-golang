package main

import "fmt"

func main() {
	sliceOfString := []string{"hari", "om", "yadav"} // slices are build on top of arrays
	fmt.Println(len(sliceOfString))

	fmt.Println(sliceOfString)
	fmt.Printf("%v", sliceOfString)
	fmt.Println()
	fmt.Println(sliceOfString[0:1])
	fmt.Println(sliceOfString[1:])
	fmt.Println(sliceOfString[:1])

	for i := 0; i < len(sliceOfString); i++ {
		fmt.Println("- ", sliceOfString[i])
	}

	for i := range sliceOfString {
		fmt.Println(i) //print only index, only key
	}

	for i, i2 := range sliceOfString { // print index and value, key and value
		fmt.Println(i," ",i2)
	}

	for _, i2 := range sliceOfString { //print only value
		fmt.Println(i2)
	}

	// add new element at the end of slice
	sliceOfString = append(sliceOfString, "chandan")
	fmt.Println(sliceOfString)


	// find max
	sliceOfInt := []int{1,22,4,1212,56}
	max := sliceOfInt[0]                // store 1st value
	for _, it := range sliceOfInt[1:] { // iterate from 2nd value to last
		if it > max {
			max = it
		}
	}
	fmt.Println(max)

}
