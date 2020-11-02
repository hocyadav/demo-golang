package main

import "fmt"

func main() {

	map1 := map[string]int{// same as slice []int -> [string]int -> map[string]int
		"hari" : 1,
		"om" : 2,
		"yadav" : 3, // comma required
	}
	fmt.Println(map1)
	fmt.Println(map1["om"])

	fmt.Println(map1["chandan"])// value not present then return 0

	value, ok := map1["chandan"]
	if !ok {
		fmt.Println("chandan not found")
	} else {
		fmt.Println(value)
	}

	// add and delete
	map1["chandan"] = 4
	fmt.Println(map1)

	// delete
	delete(map1, "yadav")
	fmt.Println(map1)

	for key := range map1 {
		fmt.Println("key - ",key)
	}
	for key, value := range map1 {
		fmt.Println("key :", key, ", value :", value)
	}

}
