package main

import "fmt"

func main() {

	var num1, num2, num3 float64

	fmt.Print("Enter num1: ")
	fmt.Scan(&num1)

	fmt.Print("Enter num2: ")
	fmt.Scan(&num2)

	fmt.Print("Enter num3: ")
	fmt.Scan(&num3)

	average := (num1 + num2 + num3) / 3

	fmt.Println("The average is ", average)

}
