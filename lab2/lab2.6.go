package main

import "fmt"

func addIntegers(a int, b int) float64 {
	return float64(a+b) / 2
}

func main() {
	var num1, num2 int

	fmt.Print("Enter the 1st number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter the 2nd number: ")
	fmt.Scan(&num2)

	result := addIntegers(num1, num2)
	fmt.Println("Average:", result)
}
