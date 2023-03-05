package learn_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

// Once adalah fitur di golang yang bisa kita gunakan untuk memastikan bahwa sebuah function dieksekusi hanya sekali
// Jadi beberapa banyak pun goroutine yang mengakses, bisa dipastikan bahwa goroutine yang pertama yang bisa mengakses function nya
//  Goroutine yang lain akan dihiraukan, artinya function tidak akan dieksekusi lagi

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter)
}
