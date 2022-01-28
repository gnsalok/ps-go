package main

import (
	"bytes"
	"fmt"
)

func main() {

	// Noob
	str1 := "String1"
	str2 := "String2"
	ans := str1 + str2
	fmt.Println(ans)

	// Pro
	var b bytes.Buffer
	b.WriteString("P")
	b.WriteString("R")
	b.WriteString("O")

	fmt.Println(b.String())

}
