package learn_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Pool adalah implementasi design pattern berdasarkan object pool pattern
// Sederhananya design pattern Pool digunakan untuk menyimpan data, selanjutnya untuk menggunakan datanya, kita bisa mengambil Pool, dan setelah selesai menggunakan datanya, kita bisa menyimpan kembali ke Pool nya
// Implementasi Pool pada golang ini sudah aman dari problem race condition

func TestPool(t *testing.T) {
	pool := sync.Pool{
		// Set default value if get pool return null
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Rizky")
	pool.Put("Alfikri")
	pool.Put("Rachmat")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(i, ".", data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Program end")
}
