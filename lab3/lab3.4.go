package main

import (
	"fmt"
)

func main() {

	var numbers [5]int

	for i := 0; i < len(numbers); i++ {
		numbers[i] = i + 1
	}

	fmt.Println("The array:", numbers)
}

/* package main

import (
	"fmt"
	"strconv"
)

func main() {

	var numbers [5]int

	for i := 0; i < len(numbers); i++ {
		var input string
		fmt.Printf("Enter the element %d: ", i+1)
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Input error. Try again.")
			i--
			continue
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Error: enter the int.")
			i--
			continue
		}

		numbers[i] = num
	}

	fmt.Println("The array:", numbers)
} */
