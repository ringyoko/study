package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumbers(ch chan int) {
	for {
		num := rand.Intn(100)
		ch <- num
		time.Sleep(500 * time.Millisecond)
	}
}

func checkEvenOdd(numbers chan int, result chan string) {
	for {
		num := <-numbers // Чтение числа из канала
		if num%2 == 0 {
			result <- fmt.Sprintf("%d is even", num)
		} else {
			result <- fmt.Sprintf("%d is odd", num)
		}
	}
}

func main() {

	numberChannel := make(chan int)
	resultChannel := make(chan string)

	go generateNumbers(numberChannel)
	go checkEvenOdd(numberChannel, resultChannel)

	// Используем select для получения данных из каналов
	for {
		select {

		case result := <-resultChannel:
			fmt.Println(result)

		case num := <-numberChannel:
			fmt.Printf("Generated number: %d\n", num)
		}
	}
}
