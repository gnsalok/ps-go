package main

import "fmt"

func main() {
	var name string
	fmt.Println("Enter your name?")
	fmt.Scanln(&name)
	print("Your name is", name)
}
