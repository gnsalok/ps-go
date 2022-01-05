package main

import "fmt"

type Bot interface {
	greeting() string
}

type EnglishBot struct{}
type SpanishBot struct{}

func (EnglishBot) greeting() string {
	return "Hi there!"
}

func (SpanishBot) greeting() string {
	return "Hola!!"
}

func printMessage(b Bot) {
	fmt.Println(b.greeting())
}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}

	printMessage(eb)
	printMessage(sb)

}
