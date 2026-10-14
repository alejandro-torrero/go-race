package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	mu      sync.Mutex // Mutual exclution
)

func increment() {
	mu.Lock()
	defer mu.Unlock()

	v := counter
	time.Sleep(time.Microsecond)

	counter = v + 1
}

func main() {
	expected := 1000
	var wg sync.WaitGroup

	for i := 0; i < expected; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()

	fmt.Println("counter:", counter)
	fmt.Println("expected:", expected)
}
