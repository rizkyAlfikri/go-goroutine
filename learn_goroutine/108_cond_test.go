package learn_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Cond adalah implemtasi locking berbasis kondisi
// Cond membutuhkan Locker (bisa menggunakan Mutex atau RWMutex) untuk implmentasi locking nya, namun berbeda dengan Locker biasanya, di Cond terdapat function Wait() untuk mengecek apakah perlu menunggu atau tidak
// Function Signal() bisa digunakan untuk memberitahu sebuah goroutine agar tidak perlu menunggu lagi. Sedangkan function Broadcast() digunkan untuk membei tahu semua goroutine agar tidak perlu menunggu lagi
// Untuk membuat Cond, kita bisa menggunakan function sync.NewCond(locker)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()

	fmt.Println("done", value)

	cond.L.Unlock()

}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	// Jalankan satu satu
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		cond.Signal()
	// 	}
	// }()

	// Jalankan sekaligus
	go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast()
	}()

	group.Wait()
}
