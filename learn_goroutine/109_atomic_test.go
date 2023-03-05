package learn_goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// Golang memiliki package yang bernamaa sync/atomic
// Atomic merupakan package yang digunakan untuk menggunakan data primitive secara aman pada process concurent
// Contoh sebelumnya kita telah menggunakan Mutex untuk melakukan process locking ketika ingin menaikan angka di counter. Hal ini sebenarnya bisa dilakukan menggunakan atomic package
func TestAtomic(t *testing.T) {
	var x int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 1000; i++ {
		go func() {
			group.Add(1)
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()

	fmt.Println("Counter = ", x)
}
