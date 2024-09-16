package main

import "fmt"

func main() {

	var number int

	fmt.Print("Enter your number: ")
	fmt.Scan(&number)

	if number%2 == 0 {
		fmt.Println("Number", number, "is even!")
	} else {
		fmt.Println("Number", number, "is odd!")
	}
}
