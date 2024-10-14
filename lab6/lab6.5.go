package main

import (
	"fmt"
	"sync"
)

type CalculationRequest struct {
	Operand1  float64
	Operand2  float64
	Operation string
	Result    chan float64
}

func calculatorServer(requests chan CalculationRequest) {
	for req := range requests {
		var result float64
		switch req.Operation {
		case "+":
			result = req.Operand1 + req.Operand2
		case "-":
			result = req.Operand1 - req.Operand2
		case "*":
			result = req.Operand1 * req.Operand2
		case "/":
			if req.Operand2 != 0 {
				result = req.Operand1 / req.Operand2
			} else {
				fmt.Println("Error: Division by zero")
				req.Result <- 0
				continue
			}
		default:
			fmt.Println("Error: Unknown operation")
			req.Result <- 0
			continue
		}

		req.Result <- result
	}
}

func client(id int, requests chan CalculationRequest, wg *sync.WaitGroup) {
	defer wg.Done()

	var operand1, operand2 float64
	var operation string

	fmt.Printf("Client %d: Enter first operand: ", id)
	fmt.Scanln(&operand1)

	fmt.Printf("Client %d: Enter operation (+, -, *, /): ", id)
	fmt.Scanln(&operation)

	fmt.Printf("Client %d: Enter second operand: ", id)
	fmt.Scanln(&operand2)

	// Канал для получения результата
	resultChannel := make(chan float64)
	req := CalculationRequest{
		Operand1:  operand1,
		Operand2:  operand2,
		Operation: operation,
		Result:    resultChannel,
	}

	fmt.Printf("Client %d: Sending request: %f %s %f\n", id, operand1, operation, operand2)
	requests <- req // Отправляем запрос на сервер калькулятора

	result := <-resultChannel
	fmt.Printf("Client %d: Result of %f %s %f = %f\n", id, operand1, operation, operand2, result)
}

func main() {
	// Канал для запросов к калькулятору
	requests := make(chan CalculationRequest)

	// Запуск серверной горутины калькулятора
	go calculatorServer(requests)

	var clientCount int
	fmt.Print("Enter number of clients: ")
	fmt.Scanln(&clientCount)

	// для ожидания завершения всех клиентов
	var wg sync.WaitGroup

	for i := 1; i <= clientCount; i++ {
		wg.Add(1)
		client(i, requests, &wg) // Запускаем клиента без горутины
	}

	wg.Wait()

	close(requests)
}
