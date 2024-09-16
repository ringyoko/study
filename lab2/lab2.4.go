package main

import "fmt"

func main() {

	var text string

	fmt.Print("Enter text: ")
	fmt.Scan(&text)

	var length = len(text)

	fmt.Print("String length: ", length)
}

/* func getStringLength(text string) int {
	return len(text)
}

func main() {

	var myInput string

	fmt.Print("Enter ur input: ")
	fmt.Scan(&myInput)

	length := getStringLength(myInput)
	fmt.Println("String length:", length)
}
*/
