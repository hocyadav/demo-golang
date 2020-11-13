package main

import "fmt"
// java class
type Student struct {
	Name string //Name means access via other packages like Public, name means private
	Age int
	Address string
}

// its like method inside class
func (receiver *Student) updateName()  {
	receiver.Name = receiver.Name + "_new"
}

// its like method inside class
func (receiver *Student) updateName2()  string{// imp add * as receiver
	receiver.Name = receiver.Name + "_new"
	return  receiver.Name
}

func main() {
	student2 := Student{// storing simple obj
		Name:    "omp",
		Age:     27,
		Address: "karnataka",
	}
	fmt.Printf("student 2 %+v \n", student2)
	student2.updateName()
	fmt.Printf("student 2 %+v \n", student2)


	student3 := &Student{ // storing as pointer - this will also work when we are calling struct method update
		Name:    "omp",
		Age:     27,
		Address: "karnataka",
	}
	fmt.Println(student3) // print pointer
	fmt.Println(*student3) // print actual value
	student3.updateName() // only update name
	fmt.Println(student3)

	name2 := student3.updateName2() // update + return value
	fmt.Println(name2)

}
