package main

import (
	"fmt"
	"strings"
)

func main() {

	map1 := map[string]int{ // same as slice []int -> [string]int -> map[string]int
		"hari":  1,
		"om":    2,
		"yadav": 3, // comma required
	}
	fmt.Println(map1)
	fmt.Println(map1["om"])

	fmt.Println(map1["chandan"]) // value not present then return 0

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
		fmt.Println("key - ", key)
	}
	for key, value := range map1 {
		fmt.Println("key :", key, ", value :", value)
	}

	text := `
hari om hari om hari chandan    chandan the hello hi bye new word
`
	fmt.Println(text)
	words := strings.Fields(text)// remve white spaces and only store words
	fmt.Println(words)
	map2 := map[string]int{} // empty map
	for _, word := range words {
		fmt.Println(word)
		toLower := strings.ToLower(word)
		map2[toLower]++
		//map2[strings.ToLower(word)]++ // write in single line
	}
	fmt.Println(map2)

}
