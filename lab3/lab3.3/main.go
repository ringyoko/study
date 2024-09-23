package main

import (
	"bufio"
	"fmt"
	"os"
	"study/stringutils"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your string: ")
	input, _ := reader.ReadString('\n')

	input = input[:len(input)-1]

	fmt.Println("Reversed string:", stringutils.ReverseString(input))
}
