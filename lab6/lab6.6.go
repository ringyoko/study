package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

// Воркеры будут получать задачи из этого канала
type Task struct {
	ID    int
	Input string
}

// Воркеры будут отправлять результаты в этот канал
type Result struct {
	ID     int
	Output string
}

func worker(id int, tasks <-chan Task, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {

		reversed := reverseString(task.Input)
		fmt.Printf("Worker %d: Task %d completed, reversed string: %s\n", id, task.ID, reversed)

		results <- Result{
			ID:     task.ID,
			Output: reversed,
		}
	}
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	/*
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			return
		}
		fmt.Println("Current working directory:", wd) */

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var tasks []Task
	scanner := bufio.NewScanner(file)
	taskID := 1
	for scanner.Scan() {
		line := scanner.Text()
		tasks = append(tasks, Task{
			ID:    taskID,
			Input: line,
		})
		taskID++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var workerCount int
	fmt.Print("Enter number of workers: ")
	fmt.Scanln(&workerCount)

	// Каналы для передачи задач и получения результатов
	taskChannel := make(chan Task, len(tasks))
	resultChannel := make(chan Result, len(tasks))

	// Группа для синхронизации завершения всех воркеров
	var wg sync.WaitGroup

	// Запуск воркеров
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go worker(i, taskChannel, resultChannel, &wg)
	}

	for _, task := range tasks {
		taskChannel <- task
	}
	close(taskChannel) // Закрываем канал задач, чтобы воркеры знали, что больше задач не будет

	// Ожидаем завершения всех воркеров
	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for result := range resultChannel {
		outputLine := fmt.Sprintf("Task %d: %s\n", result.ID, result.Output)
		writer.WriteString(outputLine)
		fmt.Printf("Result written: %s", outputLine)
	}
	writer.Flush()

	fmt.Println("All tasks processed and results written to output.txt")
}
