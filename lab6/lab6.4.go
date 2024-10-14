package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int // Общая переменная-счётчик
	mu      sync.Mutex
)

func increment(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	for i := 0; i < 100; i++ {
		// включить/выключить мьютекс
		mu.Lock()
		counter++
		mu.Unlock()

		time.Sleep(time.Millisecond * 10)
	}
}

func main() {
	var wg sync.WaitGroup
	goroutines := 5

	for i := 1; i <= goroutines; i++ {
		wg.Add(1)
		go increment(&wg, i)
	}

	wg.Wait()

	fmt.Printf("Final counter value: %d\n", counter)
}
