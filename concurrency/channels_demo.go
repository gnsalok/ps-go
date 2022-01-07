package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Channels Demo...")

	wg.Add(2)

	ch := make(chan int)

	go send(ch)

	go receive(ch)

	wg.Wait()

}

func send(ch chan<- int) {
	defer wg.Done()
	ch <- 100
}

func receive(ch <-chan int) {
	defer wg.Done()
	fmt.Println(<-ch)
}
