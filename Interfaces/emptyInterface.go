package main

import (
	"fmt"
	"reflect"
)

func printNumbers(n interface{}) {
	fmt.Println(n, reflect.TypeOf(n))

	// for i, v := range n {
	// 	fmt.Println(i, v)
	// }
}

func main() {
	lx := []int{1, 2, 3, 4, 4, 5, 5}

	printNumbers(lx)

}
