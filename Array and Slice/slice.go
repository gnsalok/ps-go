package main

import "fmt"

func main() {

	/*
		- slices are resizable array ; or Slices are Pointer/Reference to Array [V.Imp.]
		- built on the top of array.
		- slice are passed to function by reference, casue its built on the top of array.
	*/

	// Creating slices
	/*
		courses := make([]string, 5, 10)
		courses[0] = "course1"
		courses[1] = "course2"
		courses[2] = "course3"
	*/

	//Slice; Declare and initialize in one go
	courses := []string{
		"course1",
		"course2",
		"course3",
	}

	fmt.Printf("Length of slice %d \nCapacity of slice %d\n\n", len(courses), cap(courses))

	for _, c := range courses {
		fmt.Println(c)
	}

}
