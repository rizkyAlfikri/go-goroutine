package learn_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func RunHelloWordl() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWordl()
	fmt.Println("ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i <= 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 2)
	defer close(channel)

	channel <- "Rizky"
	channel <- "Alfi"

	fmt.Println(<-channel, <-channel)

	fmt.Println("End")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 20; i++ {
			channel <- "Masukan data ke " + strconv.Itoa(i)
		}

		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel1", data)
			counter++

		case data := <-channel2:
			fmt.Println("Data dari channel2", data)
			counter++
		}

		if counter >= 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel1", data)
			counter++

		case data := <-channel2:
			fmt.Println("Data dari channel2", data)
			counter++

		default:
			fmt.Println("waiting data")

		}

		if counter >= 2 {
			break
		}
	}
}
