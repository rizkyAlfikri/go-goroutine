package learn_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Input data to channel")
		channel <- "Rizky alfikri"
		fmt.Println("Finish input data to channel")
	}()

	fmt.Println("Receive data from channel")
	data := <-channel
	fmt.Println("data = ", data)

	fmt.Println("Program end")
	time.Sleep(5 * time.Second)

}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	fmt.Println("Receive data from channel")
	data := <-channel
	fmt.Println("data = ", data)

	fmt.Println("Program end")
	time.Sleep(5 * time.Second)

}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	fmt.Println("Input data to channel")
	channel <- "Rizky alfikri"
	fmt.Println("Finish input data to channel")
}
