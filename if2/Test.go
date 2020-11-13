package main

import "fmt"

func main() {

	if err := myfun(); err != nil {// if block with multiple lines
		fmt.Println(" error : ", err)
	} else {
		fmt.Println("no error..")
	}
}

func myfun() error {
	return fmt.Errorf("myfun error !!")
}
