package main

import "fmt"

func main() {

	var number int

	fmt.Print("Enter your number: ")
	fmt.Scan(&number)

	switch {

	case number > 0:
		fmt.Println("Positive.")

	case number < 0:
		fmt.Println("Negative.")

	default:
		fmt.Println("Zero.")

	}
}

/* func classifyNumber(num int) string {
	var result string

	switch {
	case num > 0:
		result = "Positive"
	case num < 0:
		result = "Negative"
	default:
		result = "Zero"
	}

	return result
}

func main() {
	var number int

	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	classification := classifyNumber(number)
	fmt.Println("The number is:", classification)
}
*/
