package learn_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

// Map ini sama seperti Map biasa, tetapi lebih safe terhadap race condition
func AddToMap(data *sync.Map, waitGroup *sync.WaitGroup, value int) {
	defer waitGroup.Done()

	waitGroup.Add(1)
	data.Store(value, value)
}

func TestSyncMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(data, group, i)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})

}
