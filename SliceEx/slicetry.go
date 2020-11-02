package main

import "fmt"

func main() {
	slice1 := []string{"hari", "om", "yadav"} // slices are build on top of arrays
	fmt.Println(len(slice1))

	fmt.Println(slice1)
	fmt.Printf("%v", slice1)
	fmt.Println()
	fmt.Println(slice1[0:1])
	fmt.Println(slice1[1:])
	fmt.Println(slice1[:1])

	for i := 0; i < len(slice1); i++ {
		fmt.Println("- ",slice1[i])
	}

	for i := range slice1 {
		fmt.Println(i) //print only index, only key
	}

	for i, i2 := range slice1 {// print index and value, key and value
		fmt.Println(i," ",i2)
	}

	for _, i2 := range slice1 {//print only value
		fmt.Println(i2)
	}

	// add new element at the end of slice
	slice1 = append(slice1, "chandan")
	fmt.Println(slice1)


	// find max
	slice2 := []int{1,22,4,1212,56}
	max := slice2[0] // store 1st value
	for _, it := range slice2[1:] {// iterate from 2nd value to last
		if it > max {
			max = it
		}
	}
	fmt.Println(max)

}
