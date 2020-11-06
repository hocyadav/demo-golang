package main

import "fmt"

type Student struct {
	Name string //Name means access via other packages like Public, name means private
	Age int
	Address string
}

func main() {
	student1 := Student{"hari", 30, "bangalore"}
	fmt.Println(student1)
	fmt.Printf("student %+v \n", student1) // +v gives nice view of struct

	student2 := Student{
		Name:    "omp",
		Age:     27,
		Address: "karnataka",
	}
	fmt.Printf("student 2 %+v \n", student2)

	student3 := Student{} //assign default values
	fmt.Printf("student 3 %+v \n", student3)

}
