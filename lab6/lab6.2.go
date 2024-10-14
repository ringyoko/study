package main

import (
	"fmt"
)

func generateFibonacci(n int, ch chan int) {

	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a // Отправляем число в канал
		a, b = b, a+b
	}
	close(ch) // Закрываем канал после завершения отправки данных
}

func printFromChannel(ch chan int) {
	for num := range ch {
		fmt.Println(num)
	}
}

func main() {

	ch := make(chan int)

	go generateFibonacci(10, ch)

	printFromChannel(ch)
}
