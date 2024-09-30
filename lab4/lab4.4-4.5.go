package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumOfNumbers() float32 {
	fmt.Print("\nPlease enter some numbers using SPACEBAR: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)
	numbers := strings.Split(input, " ")

	var sum float32
	for _, num := range numbers {

		if n, err := strconv.ParseFloat(num, 32); err == nil {
			sum += float32(n)
		} else {
			fmt.Printf("'%s' is NOT a number, dummy. I'll let this slide for now.\n", num)
		}
	}
	return sum
}

func main() {

	fmt.Print("\nHi! My creator taught me how to exclaim your teeny-tiny words and do a simple calculation. Let me show this off.\n")

	fmt.Print("Input, please: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)
	upperInput := strings.ToUpper(input)

	fmt.Println("Look, I've enlarged your letters for you ^_^ :", upperInput)

	sum := sumOfNumbers()

	fmt.Printf("I've calculated this for you ^_^ : %.2f\n", sum)

	fmt.Print("\nSee ya, I guess.")
}
