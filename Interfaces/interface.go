package main

import "fmt"

func printNumbers(n ...int) {
	for i, v := range n {
		fmt.Println(i, v)
	}

}

func main() {
	// lx := []int{1, 2, 3, 4, 4, 5, 5}
	printNumbers(1, 2, 3, 4)

}
