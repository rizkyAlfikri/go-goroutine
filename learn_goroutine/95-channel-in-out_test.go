package learn_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelInOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	fmt.Println("Program end")
	time.Sleep(5 * time.Second)
}

// use channel only to receive data
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	fmt.Println("Input data to channel")
	channel <- "Rizky alfikri"
	fmt.Println("Finish input data to channel")
}

// use channel only to receive data
func OnlyOut(channel <-chan string) {
	fmt.Println("Receive data from channel")
	data := <-channel
	fmt.Println("data = ", data)
}
