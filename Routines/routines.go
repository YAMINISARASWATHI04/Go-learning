package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
	// close(doneChan)
}
func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello !", phrase)
	// Arrow operator indicates where data should flow
	doneChan <- true
	close(doneChan)
}
func main() {
	// dones := make([]chan bool, 4)
	done := make(chan bool, 4)
	// dones[0] = make(chan bool)
	go greet("Nice to meet you!", done)
	// dones[1] = make(chan bool)
	go greet("How are you?", done)
	// dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	// dones[3] = make(chan bool)
	go greet("Bye", done)
	// After adding go keyword it will not output any result
	// It willl dispatch the four functions

	for range done {

	}
	// Waiting for data to come out of the channel using this operations
}
