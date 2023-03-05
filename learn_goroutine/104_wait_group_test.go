package learn_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsyncronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

// WaitGroup adalah fitur yang bisa digunakan untuk menunggu sebuah process selesai dilakukan
// Hal ini diperlukan, misal kita ingin menjalankan beberapa process menggunakan goroutine, tapi kita ingin semua process goroutine selesai terlebih dahulu sebelum aplikasi selesai
// Kasus seperti ini bisa menggunakan WaitGroup

func TestWaitGroup(t *testing.T) {
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsyncronous(&group)
	}

	group.Wait()
	fmt.Println("Process finished")
}
