package main

import "fmt"

func main() {

	var num1, num2 float32

	fmt.Print("Enter num1: ")
	fmt.Scan(&num1)

	fmt.Print("Enter num2: ")
	fmt.Scan(&num2)

	sum := num1 + num2
	difference := num1 - num2

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
}
