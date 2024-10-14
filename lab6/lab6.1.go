package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func factorial(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Printf("Factorial of %d: %d\n", n, result)
}

func generateRandomNumbers(count int, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		num := rand.Intn(100)
		fmt.Printf("Random number: %d\n", num)
		time.Sleep(150 * time.Millisecond)
	}
}

func sumSeries(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("Sum of the row until %d: %d\n", n, sum)
}

func main() {
	var wg sync.WaitGroup

	// Добавляем три задачи в WaitGroup
	wg.Add(3)

	go factorial(5, &wg)
	go generateRandomNumbers(5, &wg)
	go sumSeries(5, &wg)

	// Ожидание завершения всех горутин
	wg.Wait()
	fmt.Println("All complete!")
}
