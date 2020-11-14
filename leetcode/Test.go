package main

import (
	"fmt"
)

func main() {
	lengthOfLongestSubstring("abc")
}

func lengthOfLongestSubstring(s string) int {
	//map[mySetType]bool this will be my set in golang, make will create a new object
	set := make(map[byte]bool) //this is set.. no set in golan use map as set https://stackoverflow.com/questions/34018908/golang-why-dont-we-have-a-set-datastructure
	fmt.Printf("initail set/map set %v\n mset : ", set)
	l := 0
	result := 0
	for r := 0; r < len(s); r++ {
		for set[s[r]] == true { //in go no while loop, we can use for loop to make while https://tour.golang.org/flowcontrol/3
			delete(set, s[l])
			l++
		}
		charAt := s[r]              // this is char At , get index char value from string
		set[charAt] = true          // add key value in map
		result = max(result, r-l+1) // no max function in go lang we can use our own max
	}
	fmt.Println("result :", result)
	return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
