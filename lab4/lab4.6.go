package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func reverseIntArray() {
	fmt.Print("\nPlease enter some int numbers using SPACEBAR: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	input = strings.TrimSpace(input)
	numbers := strings.Split(input, " ")

	var intArray []int
	for _, num := range numbers {

		if n, err := strconv.Atoi(num); err == nil {
			intArray = append(intArray, n)
		} else {
			fmt.Printf("Jesus christ, '%s' is NOT an int!!!.. *sigh* I guess I'll let it slide. \n", num)
		}
	}

	fmt.Print("Reverse: ")
	for i := len(intArray) - 1; i >= 0; i-- {
		fmt.Printf("%d ", intArray[i])
	}
	fmt.Println()
}

func main() {

	reverseIntArray()
}
